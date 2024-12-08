package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/repo"
)

type MessageVote struct {
	repo.Repo[model.MessageVote]
	VoteCache *cache.VoteCache
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}

func NewMessageVote(db *gorm.DB, voteCache *cache.VoteCache) *MessageVote {
	return &MessageVote{Repo: repo.NewRepo[model.MessageVote](db), VoteCache: voteCache}
}

func (m *MessageVote) GetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	if uids, err := m.VoteCache.GetVoteAnswerUser(ctx, vid); err == nil {
		return uids, nil
	}

	uids, err := m.SetVoteAnswerUser(ctx, vid)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

func (m *MessageVote) SetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	uids := make([]int, 0)
	err := m.Repo.Db.WithContext(ctx).Table("message_vote_answers").Where("vote_id = ?", vid).Pluck("user_id", &uids).Error
	if err != nil {
		return nil, err
	}

	_ = m.VoteCache.SetVoteAnswerUser(ctx, vid, uids)
	return uids, nil
}

func (m *MessageVote) GetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	value, err := m.VoteCache.GetVoteStatistics(ctx, vid)
	if err != nil {
		return m.SetVoteStatistics(ctx, vid)
	}

	statistic := &VoteStatistics{}
	_ = jsonutil.Decode(value, statistic)
	return statistic, nil
}

func (m *MessageVote) SetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	var (
		vote         model.MessageVote
		answerOption map[string]any
		options      = make([]string, 0)
	)
	tx := m.Repo.Db.WithContext(ctx)
	if err := tx.Table("message_votes").First(&vote, vid).Error; err != nil {
		return nil, err
	}

	if err := jsonutil.Decode(vote.AnswerOption, &answerOption); err != nil {
		return nil, err
	}

	if err := tx.Table("message_vote_answers").
		Where("vote_id = ?", vid).
		Pluck("option", &options).
		Error; err != nil {
		return nil, err
	}

	opts := make(map[string]int)
	for option := range answerOption {
		opts[option] = 0
	}

	for _, option := range options {
		opts[option] += 1
	}

	statistic := &VoteStatistics{
		Options: opts,
		Count:   len(options),
	}

	_ = m.VoteCache.SetVoteStatistics(ctx, vid, jsonutil.Encode(statistic))

	return statistic, nil
}

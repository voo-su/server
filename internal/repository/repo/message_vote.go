package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core"
	"voo.su/pkg/jsonutil"
)

type MessageVote struct {
	core.Repo[model.MessageVote]
	Cache *cache.Vote
}

func NewMessageVote(db *gorm.DB, cache *cache.Vote) *MessageVote {
	return &MessageVote{
		Repo:  core.NewRepo[model.MessageVote](db),
		Cache: cache,
	}
}

func (t *MessageVote) GetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	if uids, err := t.Cache.GetVoteAnswerUser(ctx, vid); err == nil {
		return uids, nil
	}

	uids, err := t.SetVoteAnswerUser(ctx, vid)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

func (t *MessageVote) SetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	uids := make([]int, 0)
	err := t.Repo.Db.WithContext(ctx).Table("message_vote_answers").Where("vote_id = ?", vid).Pluck("user_id", &uids).Error
	if err != nil {
		return nil, err
	}

	_ = t.Cache.SetVoteAnswerUser(ctx, vid, uids)
	return uids, nil
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}

func (t *MessageVote) GetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	value, err := t.Cache.GetVoteStatistics(ctx, vid)
	if err != nil {
		return t.SetVoteStatistics(ctx, vid)
	}

	statistic := &VoteStatistics{}
	_ = jsonutil.Decode(value, statistic)
	return statistic, nil
}

func (t *MessageVote) SetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	var (
		vote         model.MessageVote
		answerOption map[string]any
		options      = make([]string, 0)
	)
	tx := t.Repo.Db.WithContext(ctx)
	if err := tx.Table("message_votes").First(&vote, vid).Error; err != nil {
		return nil, err
	}

	if err := jsonutil.Decode(vote.AnswerOption, &answerOption); err != nil {
		return nil, err
	}

	err := tx.Table("message_vote_answers").Where("vote_id = ?", vid).Pluck("option", &options).Error
	if err != nil {
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

	_ = t.Cache.SetVoteStatistics(ctx, vid, jsonutil.Encode(statistic))

	return statistic, nil
}

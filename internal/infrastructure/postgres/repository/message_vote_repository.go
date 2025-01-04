package repository

import (
	"context"
	"gorm.io/gorm"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/gormutil"
	"voo.su/pkg/jsonutil"
)

type MessageVoteRepository struct {
	gormutil.Repo[postgresModel.MessageVote]
	VoteCacheRepo *redisRepo.VoteCacheRepository
}

type VoteStatistics struct {
	Count   int            `json:"count"`
	Options map[string]int `json:"options"`
}

func NewMessageVoteRepository(
	db *gorm.DB,
	voteCacheRepo *redisRepo.VoteCacheRepository,
) *MessageVoteRepository {
	return &MessageVoteRepository{
		Repo:          gormutil.NewRepo[postgresModel.MessageVote](db),
		VoteCacheRepo: voteCacheRepo,
	}
}

func (m *MessageVoteRepository) GetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	if uids, err := m.VoteCacheRepo.GetVoteAnswerUser(ctx, vid); err == nil {
		return uids, nil
	}

	uids, err := m.SetVoteAnswerUser(ctx, vid)
	if err != nil {
		return nil, err
	}

	return uids, nil
}

func (m *MessageVoteRepository) SetVoteAnswerUser(ctx context.Context, vid int) ([]int, error) {
	uids := make([]int, 0)
	if err := m.Repo.Db.WithContext(ctx).
		Table("message_vote_answers").
		Where("vote_id = ?", vid).
		Pluck("user_id", &uids).Error; err != nil {
		return nil, err
	}

	_ = m.VoteCacheRepo.SetVoteAnswerUser(ctx, vid, uids)
	return uids, nil
}

func (m *MessageVoteRepository) GetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	value, err := m.VoteCacheRepo.GetVoteStatistics(ctx, vid)
	if err != nil {
		return m.SetVoteStatistics(ctx, vid)
	}

	statistic := &VoteStatistics{}
	_ = jsonutil.Decode(value, statistic)
	return statistic, nil
}

func (m *MessageVoteRepository) SetVoteStatistics(ctx context.Context, vid int) (*VoteStatistics, error) {
	var (
		vote         postgresModel.MessageVote
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

	_ = m.VoteCacheRepo.SetVoteStatistics(ctx, vid, jsonutil.Encode(statistic))

	return statistic, nil
}

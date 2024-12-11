package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"voo.su/pkg/jsonutil"
)

const (
	VoteUsersCache     = "dialog:vote:answer-users:%d"
	VoteStatisticCache = "dialog:vote:statistic:%d"
)

type VoteCacheRepository struct {
	Rds *redis.Client
}

func NewVoteCacheRepository(rds *redis.Client) *VoteCacheRepository {
	return &VoteCacheRepository{Rds: rds}
}

func (v *VoteCacheRepository) GetVoteAnswerUser(ctx context.Context, voteId int) ([]int, error) {
	val, err := v.Rds.Get(ctx, fmt.Sprintf(VoteUsersCache, voteId)).Result()
	if err != nil {
		return nil, err
	}

	var ids []int
	if err := jsonutil.Decode(val, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func (v *VoteCacheRepository) SetVoteAnswerUser(ctx context.Context, vid int, uids []int) error {
	return v.Rds.Set(ctx, fmt.Sprintf(VoteUsersCache, vid), jsonutil.Encode(uids), time.Hour*24).Err()
}

func (v *VoteCacheRepository) GetVoteStatistics(ctx context.Context, vid int) (string, error) {
	return v.Rds.Get(ctx, fmt.Sprintf(VoteStatisticCache, vid)).Result()
}

func (v *VoteCacheRepository) SetVoteStatistics(ctx context.Context, vid int, value string) error {
	return v.Rds.Set(ctx, fmt.Sprintf(VoteStatisticCache, vid), value, time.Hour*24).Err()
}

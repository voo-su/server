package cache

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

type Vote struct {
	Rds *redis.Client
}

func NewVote(rds *redis.Client) *Vote {
	return &Vote{Rds: rds}
}

func (v *Vote) GetVoteAnswerUser(ctx context.Context, voteId int) ([]int, error) {
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

func (v *Vote) SetVoteAnswerUser(ctx context.Context, vid int, uids []int) error {
	return v.Rds.Set(ctx, fmt.Sprintf(VoteUsersCache, vid), jsonutil.Encode(uids), time.Hour*24).Err()
}

func (v *Vote) GetVoteStatistics(ctx context.Context, vid int) (string, error) {
	return v.Rds.Get(ctx, fmt.Sprintf(VoteStatisticCache, vid)).Result()
}

func (v *Vote) SetVoteStatistics(ctx context.Context, vid int, value string) error {
	return v.Rds.Set(ctx, fmt.Sprintf(VoteStatisticCache, vid), value, time.Hour*24).Err()
}

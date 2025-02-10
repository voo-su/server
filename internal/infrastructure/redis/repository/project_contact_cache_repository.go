package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type ProjectContactCacheRepository struct {
	Rds *redis.Client
}

func NewProjectContactCacheRepository(rds *redis.Client) *ProjectContactCacheRepository {
	return &ProjectContactCacheRepository{
		Rds: rds,
	}
}

func (p *ProjectContactCacheRepository) IsContactRelation(ctx context.Context, uid, uid2 int) error {
	return p.Rds.Get(ctx, p.keyContactRelation(uid, uid2)).Err()
}

func (p *ProjectContactCacheRepository) SetContactRelation(ctx context.Context, uid, uid2 int) {
	p.Rds.SetEx(ctx, p.keyContactRelation(uid, uid2), "1", time.Hour*1)
}

func (p *ProjectContactCacheRepository) DelContactRelation(ctx context.Context, uid, uid2 int) {
	p.Rds.Del(ctx, p.keyContactRelation(uid, uid2))
}

func (p *ProjectContactCacheRepository) IsGroupRelation(ctx context.Context, uid int, gid uuid.UUID) error {
	return p.Rds.Get(ctx, p.keyGroupRelation(uid, gid)).Err()
}

func (p *ProjectContactCacheRepository) SetGroupRelation(ctx context.Context, uid int, gid uuid.UUID) {
	p.Rds.SetEx(ctx, p.keyGroupRelation(uid, gid), "1", time.Hour*1)
}

func (p *ProjectContactCacheRepository) DelGroupRelation(ctx context.Context, uid int, gid uuid.UUID) {
	p.Rds.Del(ctx, p.keyGroupRelation(uid, gid))
}

func (p *ProjectContactCacheRepository) BatchDelGroupRelation(ctx context.Context, uids []int, gid uuid.UUID) {
	for _, uid := range uids {
		p.DelGroupRelation(ctx, uid, gid)
	}
}

func (p *ProjectContactCacheRepository) keyContactRelation(uid, uid2 int) string {
	if uid2 < uid {
		uid, uid2 = uid2, uid
	}

	return fmt.Sprintf("im:contact:relation:%d_%d", uid, uid2)
}

func (p *ProjectContactCacheRepository) keyGroupRelation(uid int, gid uuid.UUID) string {
	return fmt.Sprintf("im:contact:relation:%d_%d", uid, gid)
}

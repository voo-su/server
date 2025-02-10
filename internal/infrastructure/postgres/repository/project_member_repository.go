package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type ProjectMemberRepository struct {
	gormutil.Repo[model.ProjectMember]
}

func NewProjectMemberRepository(db *gorm.DB) *ProjectMemberRepository {
	return &ProjectMemberRepository{Repo: gormutil.NewRepo[model.ProjectMember](db)}
}

func (p *ProjectMemberRepository) GetMemberIds(ctx context.Context, projectId uuid.UUID) []int {
	var ids []int
	_ = p.Repo.Model(ctx).
		Select("user_id").
		Where("project_id = ?", projectId).
		Scan(&ids)

	return ids
}

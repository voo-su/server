package repository

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/repo"
)

type ProjectMemberRepository struct {
	repo.Repo[model.ProjectMember]
}

func NewProjectMemberRepository(db *gorm.DB) *ProjectMemberRepository {
	return &ProjectMemberRepository{Repo: repo.NewRepo[model.ProjectMember](db)}
}

func (p *ProjectMemberRepository) GetMemberIds(ctx context.Context, projectId int) []int {
	var ids []int
	_ = p.Repo.Model(ctx).Select("user_id").
		Where("project_id = ?", projectId).
		Scan(&ids)

	return ids
}

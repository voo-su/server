package repo

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/pkg/repo"
)

type ProjectMember struct {
	repo.Repo[model.ProjectMember]
}

func NewProjectMember(db *gorm.DB) *ProjectMember {
	return &ProjectMember{Repo: repo.NewRepo[model.ProjectMember](db)}
}

func (p *ProjectMember) GetMemberIds(ctx context.Context, projectId int) []int {
	var ids []int
	_ = p.Repo.Model(ctx).Select("user_id").
		Where("project_id = ?", projectId).
		Scan(&ids)

	return ids
}

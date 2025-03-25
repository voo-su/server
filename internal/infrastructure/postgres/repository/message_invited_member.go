package repository

import (
	"gorm.io/gorm"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/gormutil"
)

type MessageInvitedMemberRepository struct {
	gormutil.Repo[model.MessageInvitedMember]
}

func NewMessageInvitedMemberForwardedRepository(db *gorm.DB) *MessageInvitedMemberRepository {
	return &MessageInvitedMemberRepository{Repo: gormutil.NewRepo[model.MessageInvitedMember](db)}
}

package postgres

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/postgres/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewUserSessionRepository,
	repository.NewContactRepository,
	repository.NewContactFolderRepository,
	repository.NewGroupMemberRepository,
	repository.NewUserRepository,
	repository.NewGroupChatRepository,
	repository.NewGroupChatApplyRepository,
	repository.NewGroupChatAdsRepository,
	repository.NewChatRepository,
	repository.NewMessageRepository,
	repository.NewMessageVoteRepository,
	repository.NewMessageForwardedForwardedRepository,
	repository.NewMessageAuthLogForwardedRepository,
	repository.NewMessageInvitedMemberForwardedRepository,
	repository.NewMessageMediaForwardedRepository,
	repository.NewFileSplitRepository,
	repository.NewSequenceRepository,
	repository.NewBotRepository,
	repository.NewStickerRepository,
	repository.NewPushTokenRepository,
	repository.NewProjectRepository,
	repository.NewProjectMemberRepository,
	repository.NewProjectTaskTypeRepository,
	repository.NewProjectTaskRepository,
	repository.NewProjectTaskCommentRepository,
	repository.NewProjectTaskCoexecutorRepository,
	repository.NewProjectTaskWatcherRepository,
	repository.NewFileRepository,
)

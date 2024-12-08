package repository

import (
	"github.com/google/wire"
	"voo.su/internal/repository/repo"
)

var ProviderSet = wire.NewSet(
	NewSource,
	repo.NewUserSession,
	repo.NewContact,
	repo.NewContactFolder,
	repo.NewGroupMember,
	repo.NewUser,
	repo.NewGroupChat,
	repo.NewGroupChatApply,
	repo.NewGroupChatAds,
	repo.NewChat,
	repo.NewMessage,
	repo.NewMessageVote,
	repo.NewFileSplit,
	repo.NewSequence,
	repo.NewBot,
	repo.NewSticker,
	repo.NewPushToken,
	repo.NewProject,
	repo.NewProjectMember,
	repo.NewProjectTaskType,
	repo.NewProjectTask,
	repo.NewProjectTaskComment,
	repo.NewProjectTaskCoexecutor,
	repo.NewProjectTaskWatcher,
)

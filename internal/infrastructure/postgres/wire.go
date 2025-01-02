// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
)

// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package redis

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/redis/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewCaptchaCacheRepository,
	repository.NewClientCacheRepository,
	repository.NewContactRemarkCacheRepository,
	repository.NewRedisLockCacheRepository,
	repository.NewMessageCacheRepository,
	repository.NewRelationCacheRepository,
	repository.NewRoomCacheRepository,
	repository.NewSequenceCacheRepository,
	repository.NewJwtTokenCacheRepository,
	repository.NewServerCacheRepository,
	repository.NewSmsCacheRepository,
	repository.NewVoteCacheRepository,
	repository.NewUnreadCacheRepository,
	repository.NewGroupChatRequestCacheRepository,
)

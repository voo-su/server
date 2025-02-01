package redis

import (
	"github.com/google/wire"
	"voo.su/internal/infrastructure/redis/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewCaptchaCacheRepository,
	repository.NewClientCacheRepository,
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

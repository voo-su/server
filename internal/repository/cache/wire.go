package cache

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewCaptchaCache,
	NewClientCache,
	NewContactRemarkCache,
	NewRedisLockCache,
	NewMessageCache,
	NewRelationCache,
	NewRoomCache,
	NewSequenceCache,
	NewJwtTokenCache,
	NewServerCache,
	NewSmsCache,
	NewVoteCache,
	NewUnreadCache,
	NewGroupChatRequestCache,
)

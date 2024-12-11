package process

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewServer,
	NewHealthSubscribe,
	NewMessageSubscribe,
)

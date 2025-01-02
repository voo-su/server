// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package process

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewServer,
	NewHealthSubscribe,
	NewMessageSubscribe,
)

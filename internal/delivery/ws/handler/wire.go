// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package handler

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Struct(new(ChatChannel), "*"),
)

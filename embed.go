// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package voo_su

import "embed"

//go:embed migrations
var migration embed.FS

func Migration() embed.FS {
	return migration
}

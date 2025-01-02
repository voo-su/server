// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package resource

import "embed"

//go:embed "template"
var template embed.FS

func Template() embed.FS {
	return template
}

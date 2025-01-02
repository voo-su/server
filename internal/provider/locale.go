// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package provider

import (
	"voo.su/pkg/locale"
)

func NewLocale() locale.ILocale {
	return locale.NewLocale([]string{
		"internal/locale/ru.json",
		"internal/locale/en.json",
	})
}

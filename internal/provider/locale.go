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

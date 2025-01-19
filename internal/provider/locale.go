package provider

import (
	vooSu "voo.su"
	"voo.su/internal/config"
	"voo.su/pkg/locale"
)

func NewLocale(conf *config.Config) locale.ILocale {
	return locale.NewLocale(
		vooSu.Locale(),
		[]string{"ru", "en"},
		conf.App.GetDefaultLang(),
	)
}

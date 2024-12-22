package locale

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type ILocale interface {
	Localize(key string) string
}

var _ ILocale = (*Locale)(nil)

type Locale struct {
	Bundle *i18n.Bundle
}

func NewLocale(paths []string) *Locale {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, path := range paths {
		_, err := bundle.LoadMessageFile(path)
		if err != nil {
			break
		}
	}

	return &Locale{Bundle: bundle}
}

func (l *Locale) Localize(key string) string {
	local := i18n.NewLocalizer(l.Bundle, "ru")
	message, err := local.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		log.Println("Ошибка при локализации:", err)
		return key
	}

	return message
}

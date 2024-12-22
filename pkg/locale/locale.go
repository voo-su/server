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
	bundle *i18n.Bundle
}

func NewLocale() *Locale {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFile("internal/locale/ru.json")
	if err != nil {
		log.Fatal(err)
	}

	return &Locale{bundle: bundle}
}

func (l *Locale) Localize(key string) string {
	local := i18n.NewLocalizer(l.bundle, "ru")
	message, err := local.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		log.Println("Ошибка при локализации:", err)
		return key
	}

	return message
}

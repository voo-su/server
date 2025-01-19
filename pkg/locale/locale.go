package locale

import (
	"embed"
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type ILocale interface {
	Localize(key string) string

	SetFromHeaderAcceptLanguage(acceptLang string)
}

var _ ILocale = (*Locale)(nil)

type Locale struct {
	Bundle      *i18n.Bundle
	Local       string
	DefaultLang string
}

func NewLocale(fs embed.FS, langs []string, defaultLang string) *Locale {
	bundle := i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, lang := range langs {
		filePath := "internal/locale/" + lang + ".json"
		data, err := fs.ReadFile(filePath)
		if err != nil {
			log.Printf("Error reading localization file '%s': %v", filePath, err)
			continue
		}

		if _, err := bundle.ParseMessageFileBytes(data, filePath); err != nil {
			log.Printf("Error parsing localization file '%s': %v", filePath, err)
		}
	}

	return &Locale{
		Bundle:      bundle,
		Local:       language.Russian.String(),
		DefaultLang: defaultLang,
	}
}

func (l *Locale) Localize(key string) string {
	local := i18n.NewLocalizer(l.Bundle, l.Local)
	message, err := local.Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		log.Printf("Error localizing key '%s': %v", key, err)
		return key
	}

	return message
}

func (l *Locale) SetFromHeaderAcceptLanguage(acceptLang string) {
	if acceptLang == "" {
		log.Println("Accept-Language header is empty, using default language:", l.DefaultLang)
		l.Local = l.DefaultLang
		return
	}

	matcher := language.NewMatcher(l.Bundle.LanguageTags())
	tags, _, _ := language.ParseAcceptLanguage(acceptLang)
	tag, _, confidence := matcher.Match(tags...)

	if confidence == language.No {
		log.Printf("Unknown language, using default language: %s", l.DefaultLang)
		l.Local = l.DefaultLang
		return
	}

	l.Local = tag.String()
}

// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ruTranslations "github.com/go-playground/validator/v10/translations/ru"
	"reflect"
)

var trans ut.Translator

func init() {
	_ = Initialize()
}

func Initialize() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		chinese := ru.New()
		uni := ut.New(chinese)
		trans, _ = uni.GetTranslator("ru")
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		registerCustomValidator(v, trans)

		return ruTranslations.RegisterDefaultTranslations(v, trans)
	}
	return nil
}

func Translate(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			return err.Translate(trans)
		}
	}

	return err.Error()
}

func registerCustomValidator(v *validator.Validate, trans ut.Translator) {
	phone(v, trans)
	ids(v, trans)
}

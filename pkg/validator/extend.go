package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func phone(v *validator.Validate, trans ut.Translator) {
	_ = v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		matched, _ := regexp.MatchString("^1[3456789][0-9]{9}$", fl.Field().String())
		return matched
	})

	_ = v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "Неверный формат номера телефона", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field(), fe.Field())
		return t
	})
}

func ids(v *validator.Validate, trans ut.Translator) {
	_ = v.RegisterValidation("ids", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return true
		}

		matched, _ := regexp.MatchString("^\\d+(\\,\\d+)*$", value)
		return matched
	})

	_ = v.RegisterTranslation("ids", trans, func(ut ut.Translator) error {
		return ut.Add("ids", "Неверный формат ids", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ids", fe.Field(), fe.Field())
		return t
	})
}

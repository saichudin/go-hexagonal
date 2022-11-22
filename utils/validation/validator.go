package validation

import (
	"errors"

	"github.com/go-playground/locales/id"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	id := id.New()
	uni := translator.New(id, id)

	//---split into function---
	// translate into bahasa
	var ok bool
	trans, ok := uni.GetTranslator("id")
	if !ok {
		return errors.New("cannot find translate")
	}

	err := id_translations.RegisterDefaultTranslations(cv.Validator, trans)
	if err != nil {
		return errors.New("cannot register translation")
	}
	//---split into function---

	return cv.Validator.Struct(i)
}

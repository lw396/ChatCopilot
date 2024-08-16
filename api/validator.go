package api

import (
	"github.com/lw396/ChatCopilot/internal/errors"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	v     *validator.Validate
	trans ut.Translator
}

func NewValidator() *Validator {
	v := validator.New()
	uni := ut.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(v, trans)
	return &Validator{
		v:     v,
		trans: trans,
	}
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.v.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New(errors.CodeInvalidParam, "invalid params")
		}

		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.New(errors.CodeInvalidParam, ve[0].Translate(v.trans))
		}

		return errors.New(errors.CodeInvalidParam, err.Error())
	}
	return nil
}

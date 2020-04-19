package main

import (
	"context"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gorilla/schema"
	"github.com/pdeguing/empire-and-foundation/data"
	"github.com/pdeguing/empire-and-foundation/ent/user"
	"reflect"
	"strings"
)

var decoder *schema.Decoder
var validate *validator.Validate
var translator ut.Translator

func init() {
	decoder = schema.NewDecoder()
	decoder.SetAliasTag("json")
	decoder.IgnoreUnknownKeys(true)

	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("name"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	english := en.New()
	uni := ut.New(english, english)
	translator, _ = uni.GetTranslator("en")
	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		panic(err)
	}

	err := validate.RegisterValidationCtx("unique_username", func(ctx context.Context, fl validator.FieldLevel) bool {
		n, err := data.Client.User.
			Query().
			Where(user.UsernameEQ(fl.Field().String())).
			Count(ctx)
		return err != nil && n == 0
	}, false)
	if err != nil {
		panic(err)
	}
	err = validate.RegisterTranslation("unique_username", translator, func(ut ut.Translator) error {
		return ut.Add("unique_username", "this {0} is already in use", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique_username", fe.Field())
		return t
	})
	if err != nil {
		panic(err)
	}
}

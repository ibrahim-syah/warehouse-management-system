package validatorutils

import (
	"reflect"
	"strings"
	passwordutils "warehouse-management-system/utils/password"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator(en.Locale())
		en_translations.RegisterDefaultTranslations(v, trans)

		v.RegisterValidation("password", passwordStrengthCheck)
		v.RegisterTranslation(
			"password",
			trans,
			registerTranslator("password", "{0} must have at least 8 characters, 1 upper case character, 1 number, and 1 special character"),
			translate,
		)
	}
}

var trans ut.Translator

func GetTranslator() ut.Translator {
	return trans
}

var passwordStrengthCheck validator.Func = func(fl validator.FieldLevel) bool {
	sevenOrMore, number, upper, special := passwordutils.CheckPasswordStrength(fl.Field().String())
	return sevenOrMore && number && upper && special
}

func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}

		return nil
	}
}

func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}

	return msg
}

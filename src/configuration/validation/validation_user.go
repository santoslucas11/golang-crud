package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"
	"github.com/santoslucas11/goland-crud/src/configuration/rest_err"

	ut "github.com/go-playground/universal-translator"

	english_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if translate_validation, ok := binding.Validator.Engine().(*validator.Validate); ok {
		english_translate := en.New()
		universal_translator := ut.New(english_translate, english_translate)

		transl, _ = universal_translator.GetTranslator("english")

		english_translation.RegisterDefaultTranslations(translate_validation, transl)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, erro := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: erro.Translate(transl),
				Field:   erro.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}

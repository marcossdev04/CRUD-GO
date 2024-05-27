package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	rest_rr "github.com/marcossdev04/crud-go/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}
func ValidateUserError(
	validation_err error,
) *rest_rr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_rr.NewBadRequestError("Invalid field Type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_rr.Causes{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_rr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_rr.NewBadRequestValidationError("Some field are invalid", errorCauses)
	} else {
		return rest_rr.NewBadRequestError("Error trying to convert fields")
	}
}

package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
		Message     string
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func NewXValidator() *XValidator {
	return &XValidator{
		Validator: validate,
	}
}

func (v *XValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true
			elem.Message = formatErrorMessage(err)

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "min":
		return fmt.Sprintf("[%s]: '%v' | Needs to implement 'min %s'", err.Field(), err.Value(), err.Param())
	default:
		return fmt.Sprintf("[%s]: '%v' | Needs to implement '%s'", err.Field(), err.Value(), err.Tag())
	}
}

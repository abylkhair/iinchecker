package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/wildegor/kaspi-rest/pkg/validators"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	validate := validator.New()

	iinValidator := validators.NewIINValidator()

	// Custom validation for IIN fields.
	_ = validate.RegisterValidation("iin", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := iinValidator.IsValid(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfdfa
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	var returnErrs []ValidationError

	if errs, ok := v.validate.Struct(i).(validator.ValidationErrors); ok {
		for _, err := range errs {
			ve := ValidationError{err}
			returnErrs = append(returnErrs, ve)
		}
	}

	return returnErrs
}

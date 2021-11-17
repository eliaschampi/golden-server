package rules

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(structData interface{}) []*ErrorResponse {

	var errors []*ErrorResponse

	validate := validator.New()

	err := validate.Struct(structData)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var field ErrorResponse
			field.FailedField = err.StructNamespace()
			field.Tag = err.Tag()
			field.Value = err.Param()
			errors = append(errors, &field)
		}
	}

	return errors
}

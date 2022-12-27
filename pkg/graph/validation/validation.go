package validation

import (
	"fmt"

	validateErr "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/domain/error"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

func init() {
	Validate = validator.New()

}

func ValidateInputModel(m any) *gqlerror.Error {
	validationErrors, err := validateModel(m)
	if err != nil {
		return &gqlerror.Error{
			Message:    validateErr.ErrorMessage(validateErr.ValidationError),
			Extensions: validateErr.InternalServerErrorExtension(),
		}
	}
	if len(validationErrors) > 0 {
		return &gqlerror.Error{
			Message:    validateErr.ErrorMessage(validateErr.BadInput),
			Extensions: validateErr.BadInputExtension(validationErrors),
		}
	}

	return nil
}

func validateModel(model any) (map[string]string, error) {
	if err := Validate.Struct(model); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}

		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string, len(errs))
		for _, ve := range errs {
			validationErrors[ve.StructNamespace()] = msgForTag(ve)
		}

		return validationErrors, nil
	}
	return nil, nil
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "入力は必須です"
	case "len":
		return fmt.Sprintf("%s文字で入力してください", fe.Param())
	case "gte":
		return fmt.Sprintf("%s文字以上で入力してください", fe.Param())
	case "lte":
		return fmt.Sprintf("%s文字以下で入力してください", fe.Param())
	case "timezone":
		return "IANA Time Zone databaseの形式で入力してください"
	case "HH:mm":
		return "00:00 ~ 23:59の間で入力してください"
	}
	return fe.Error()
}

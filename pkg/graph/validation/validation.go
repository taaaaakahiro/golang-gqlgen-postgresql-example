package validation

import (
	"fmt"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	customErr "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/customerr"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var (
	validate *validator.Validate

	hhmmRegex = regexp.MustCompile(`^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$`)
)

func init() {
	validate = validator.New()

	// register time validation
	if err := validate.RegisterValidation("HH:mm", func(fl validator.FieldLevel) bool {
		return hhmmRegex.MatchString(fl.Field().String())
	}, false); err != nil {
		log.Fatalln("failed to register validation")
	}
}

func ValidateModel(model any) (map[string]string, error) {
	if err := validate.Struct(model); err != nil {
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

func ValidateInputModel(m any) *gqlerror.Error {
	validationErrors, err := ValidateModel(m)
	if err != nil {
		return &gqlerror.Error{
			Message:    customErr.ErrorMessage(customErr.InternalServerError),
			Extensions: customErr.InternalServerErrorExtension(),
		}
	}
	if len(validationErrors) > 0 {
		return &gqlerror.Error{
			Message:    customErr.ErrorMessage(customErr.BadUserInput),
			Extensions: customErr.BadUserInputExtension(validationErrors),
		}
	}
	return nil
}

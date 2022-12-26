package resolver

import (
	"log"

	validateErr "github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/error"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func validateInputModel(m any) *gqlerror.Error {
	validationErrors, err := validation.ValidateModel(m)
	log.Println(validationErrors, err)
	if err != nil {
		return &gqlerror.Error{
			Message:    validateErr.ErrorMessage(validateErr.ValidationError),
			Extensions: validateErr.InternalServerErrorExtension(),
		}
	}
	if len(validationErrors) > 0 {
		return &gqlerror.Error{
			Message:    validateErr.ErrorMessage(validateErr.BadInput),
			Extensions: validateErr.BadUserInputExtension(validationErrors),
		}
	}

	return nil
}

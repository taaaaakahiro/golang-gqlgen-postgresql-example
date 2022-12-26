package resolver

import (
	"log"

	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/customer"
	"github.com/taaaaakahiro/golang-gqlgen-postgresql-example/pkg/graph/validation"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func validateInputModel(m any) *gqlerror.Error {
	validationErrors, err := validation.ValidateModel(m)
	if err != nil {
		log.Fatal(err)
		return &gqlerror.Error{
			Message:    customer.ErrorMessage(customer.InternalServerError),
			Extensions: customer.InternalServerErrorExtension(),
		}
	}
	if len(validationErrors) > 0 {
		return &gqlerror.Error{
			Message:    customer.ErrorMessage(customer.BadUserInput),
			Extensions: customer.BadUserInputExtension(validationErrors),
		}
	}

	return nil
}

package error

import "github.com/vektah/gqlparser/v2/gqlerror"

func GQL404UserError() error {
	return &gqlerror.Error{
		Message: "user is not found",
		Extensions: map[string]interface{}{
			"status": 404,
		},
	}
}

func GQL500UserError() error {
	return &gqlerror.Error{
		Message: "failed to get user",
		Extensions: map[string]interface{}{
			"status": 500,
		},
	}
}

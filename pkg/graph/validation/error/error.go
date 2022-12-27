package error

type errorCode string

const (
	defaultErrorMsg string = "internal server error"

	BadInput    errorCode = "BAD_INPUT"
	BadInputMsg string    = "invalid input"

	ValidationError    errorCode = "VALIDATE_ERROR"
	ValidationErrorMsg string    = "validate error"
)

func ErrorMessage(e errorCode) string {
	switch e {
	case BadInput:
		return BadInputMsg
	case ValidationError:
		return ValidationErrorMsg
	default:
		return defaultErrorMsg
	}
}

func BadUserInputExtension(invalidFields map[string]string) map[string]any {
	return map[string]any{
		"code":          BadInput,
		"invalidFields": invalidFields,
	}
}

func InternalServerErrorExtension() map[string]any {
	return map[string]any{
		"code": ValidationError,
	}
}

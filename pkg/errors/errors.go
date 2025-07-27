package errors

import (
	goerrors "errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Value   any
	Message string
	Err     error
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("doğrulama hatası: alan '%s' geçersiz '%v'. Neden: %s", e.Field, e.Value, e.Message)
	}
	return fmt.Sprintf("doğrulama hatası: %s", e.Message)
}

func (e *ValidationError) Unwrap() error {
	return e.Err
}

func NewValidationError(field string, value any, err error) error {
	message := ""
	if err != nil {
		message = err.Error()
	}

	return &ValidationError{
		Field:   field,
		Value:   value,
		Message: message,
		Err:     err,
	}
}

var (
	ErrInvalidInput = goerrors.New("geçersiz girdi")
)

func IsValidationError(err error) (*ValidationError, bool) {
	var validationErr *ValidationError
	if goerrors.As(err, &validationErr) {
		return validationErr, true
	}
	return nil, false
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

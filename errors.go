package nerrors

import (
	"github.com/pkg/errors"
)

type ErrorType int

const (
	NoType = ErrorType(iota)
	Internal
	BadRequest
	Validation
	PermissionDenied
)

type customError struct {
	errorType     ErrorType
	originalError error
}

func (err *customError) Error() string {
	return err.originalError.Error()
}

func (errType ErrorType) New(msg string) error {
	return &customError{
		errorType:     errType,
		originalError: errors.New(msg),
	}
}

func (errType ErrorType) Wrap(err error, msg string) error {
	return &customError{errorType: errType, originalError: errors.Wrapf(err, msg)}
}

func GetType(err error) ErrorType {
	if customErr, ok := err.(*customError); ok {
		return customErr.errorType
	}
	return NoType
}

func GetError(err error) error {
	if customErr, ok := err.(*customError); ok {
		return customErr.originalError
	}
	return err
}

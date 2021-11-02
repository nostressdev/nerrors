package nerrors

import (
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorType int

const (
	NoType = ErrorType(iota)
	Internal
	BadRequest
	Validation
	PermissionDenied
	NotFound
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

func (errType ErrorType) Newf(format string, args ...interface{}) error {
	return &customError{
		errorType:     errType,
		originalError: fmt.Errorf(format, args...),
	}
}

func (errType ErrorType) Wrap(err error, msg string) error {
	return &customError{errorType: errType, originalError: errors.Wrapf(err, msg)}
}

func (errType ErrorType) Wrapf(err error, format string, args ...interface{}) error {
	return &customError{errorType: errType, originalError: errors.Wrapf(err, format, args...)}
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

var errCodeToGRPC = map[ErrorType]codes.Code{
	NoType:           codes.Unknown,
	Internal:         codes.Internal,
	Validation:       codes.InvalidArgument,
	PermissionDenied: codes.PermissionDenied,
	NotFound:         codes.NotFound,
	BadRequest:       codes.InvalidArgument,
}

func GetErrorGRPC(err error) error {
	if status.Code(err) != codes.Unknown {
		return err
	}
	return status.Error(errCodeToGRPC[GetType(err)], err.Error())
}

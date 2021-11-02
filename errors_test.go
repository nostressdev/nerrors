package nerrors

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestErrorType_New(t *testing.T) {
	const errText = "test error"
	err := Internal.New(errText)
	assert.Equal(t, errText, err.Error())
	assert.Equal(t, Internal, GetType(err))
}

func TestErrorType_Wrap(t *testing.T) {
	const WrapText = "test"
	originalErr := errors.New("original error")
	err := BadRequest.Wrap(originalErr, WrapText)
	assert.Equal(t, BadRequest, GetType(err))
	assert.Equal(t, fmt.Sprintf("%v: %v", WrapText, originalErr), err.Error())
	assert.Equal(t, true, errors.Is(GetError(err), originalErr))
}

func TestErrorType_Newf(t *testing.T) {
	err := Internal.Newf("%s", "xxx")
	assert.Equal(t, "xxx", err.Error())
	assert.Equal(t, Internal, GetType(err))
}

func TestErrorType_Wrapf(t *testing.T) {
	const WrapText = "test"
	originalErr := errors.New("original error")
	err := BadRequest.Wrapf(originalErr, "%s", WrapText)
	assert.Equal(t, BadRequest, GetType(err))
	assert.Equal(t, fmt.Sprintf("%v: %v", WrapText, originalErr), err.Error())
	assert.Equal(t, true, errors.Is(GetError(err), originalErr))
}

func TestGrpcConvert(t *testing.T) {
	assert.Equal(t, codes.Internal, status.Code(GetErrorGRPC(Internal.New(""))))
	assert.Equal(t, codes.InvalidArgument, status.Code(GetErrorGRPC(BadRequest.New(""))))
	assert.Equal(t, codes.InvalidArgument, status.Code(GetErrorGRPC(Validation.New(""))))
	assert.Equal(t, codes.NotFound, status.Code(GetErrorGRPC(NotFound.New(""))))
	assert.Equal(t, codes.PermissionDenied, status.Code(GetErrorGRPC(PermissionDenied.New(""))))
}

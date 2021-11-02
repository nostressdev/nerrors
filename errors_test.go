package nerrors

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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

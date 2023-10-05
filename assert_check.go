package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type assertCheck struct {
	t             *testing.T
	expectedValue any
}

func (a *assertCheck) Equal(actual any) bool {
	return assert.Equal(a.t, a.expectedValue, actual)
}

func (a *assertCheck) ErrorIs(target error) bool {
	expectedErr, ok := a.expectedValue.(error)
	if !ok {
		panic("not error type")
	}

	return assert.ErrorIs(a.t, expectedErr, target)
}

func (a *assertCheck) ErrorIsNot(target error) bool {
	expectedErr, ok := a.expectedValue.(error)
	if !ok {
		panic("not error type")
	}

	return assert.NotErrorIs(a.t, expectedErr, target)
}

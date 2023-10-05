package assert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type assertSetup struct {
	t *testing.T
}

func (a *assertSetup) AssertThat(expected any) *assertCheck {
	return &assertCheck{
		t:             a.t,
		expectedValue: expected,
	}
}

func (a *assertSetup) True(expected bool, msgAndArgs ...any) bool {
	return assert.True(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) False(expected bool, msgAndArgs ...any) bool {
	return assert.False(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) NotNil(expected any, msgAndArgs ...any) bool {
	return assert.NotNil(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) Nil(expected any, msgAndArgs ...any) bool {
	return assert.Nil(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) Error(expected error, msgAndArgs ...any) bool {
	return assert.Error(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) NoError(expected error, msgAndArgs ...any) bool {
	return assert.NoError(a.t, expected, msgAndArgs...)
}

func (a *assertSetup) CatchPanic(runnable func()) (ap *assertPanic) {
	defer func() {
		ap = &assertPanic{
			rec: recover(),
			t:   a.t,
		}
	}()

	runnable()

	return ap
}

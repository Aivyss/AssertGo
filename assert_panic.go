package assert

import (
	"errors"
	errors2 "github.com/aivyss/AssertGo/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type assertPanic struct {
	rec any
	t   *testing.T
}

func (a *assertPanic) IsError(msgAndARgs ...interface{}) *assertPanic {
	err := a.convertToErr(msgAndARgs)

	assert.Error(a.t, err)

	return a
}

func (a *assertPanic) convertToErr(msgAndARgs ...interface{}) error {
	if a.rec == nil {
		var err error = errors2.MsgAndArgsErr{
			Err:        errors.New("no one is caught"),
			MsgAndArgs: msgAndARgs,
		}

		panic(err)
	}

	err, ok := a.rec.(error)

	if !ok {
		var err error = errors2.MsgAndArgsErr{
			Err:        errors.New("the caught value is not a error"),
			MsgAndArgs: msgAndARgs,
		}

		panic(err)
	}

	return err
}

func (a *assertPanic) ErrorMsgIs(msg string) *assertPanic {
	err := a.convertToErr()

	assert.Equal(a.t, msg, err.Error())

	return a
}

func (a *assertPanic) ErrorTypeIs(convertable func(obj interface{}) bool) *assertPanic {
	assert.True(a.t, convertable(a.rec))

	return a
}

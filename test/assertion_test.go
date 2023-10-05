package test

import (
	"errors"
	assert "github.com/aivyss/AssertGo"
	errors2 "github.com/aivyss/AssertGo/errors"
	"testing"
)

func TestAssertion(t *testing.T) {
	t.Run("AssertThat-Equal", func(t *testing.T) {
		assert.Test(t).AssertThat("abcd").Equal("abcd")
	})

	t.Run("True/False", func(t *testing.T) {
		assert.Test(t).True(true)
		assert.Test(t).False(false)
	})

	t.Run("CatchPanic-IsError/ErrorMsgIs/ErrorTypeIs", func(t *testing.T) {
		errMsg := "test error"
		panicGenerator := func() {
			panic(errors2.MsgAndArgsErr{
				Err: errors.New(errMsg),
			})
		}

		assert.Test(t).CatchPanic(func() {
			panicGenerator()
		}).IsError().ErrorMsgIs(errMsg).ErrorTypeIs(assert.ErrorTypeIs[errors2.MsgAndArgsErr])

		assert.Test(t).CatchPanic(func() {}).NoCatch()
	})
}

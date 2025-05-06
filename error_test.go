package anyutil

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyCustomError struct {
	Msg string
}

func (e *MyCustomError) Error() string {
	return e.Msg
}

type MyCustomError2 struct {
	Msg string
}

func (e *MyCustomError2) Error() string {
	return e.Msg
}

func TestAs_GenericWrapper(t *testing.T) {
	inner := &MyCustomError{Msg: "something went wrong"}
	wrapped := errors.New("wrapped: " + inner.Error())
	joined := errors.Join(inner, wrapped)

	t.Run("matches custom error", func(t *testing.T) {
		got, ok := As[*MyCustomError](joined)
		assert.True(t, ok)
		assert.NotNil(t, got)
		assert.Equal(t, "something went wrong", got.Msg)
	})

	t.Run("does not match unrelated type", func(t *testing.T) {
		_, ok := As[*MyCustomError2](joined)
		assert.False(t, ok)
	})

	t.Run("returns zero value on nil input", func(t *testing.T) {
		var err error
		got, ok := As[*MyCustomError](err)
		assert.False(t, ok)
		assert.Nil(t, got)
	})
}

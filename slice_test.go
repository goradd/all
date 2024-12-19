package all

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertSlice(t *testing.T) {
	l1 := []string{"a", "b"}
	var i1 []any
	var nilSlice []any
	i1 = ConvertSlice[any](l1)
	assert.Len(t, i1, 2)
	assert.Equal(t, "a", i1[0].(string))

	assert.Nil(t, ConvertSlice[any](nilSlice))

	var l2 []string
	assert.Nil(t, ConvertSlice[any](l2))

	i3 := []any{1, 2}
	i4 := ConvertSlice[int](i3)
	assert.True(t, i4[0] == 1)
}

func TestIsSlice(t *testing.T) {
	assert.True(t, IsSlice([]string{}))
	assert.False(t, IsSlice(5))
}

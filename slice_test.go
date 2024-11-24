package any

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopySlice(t *testing.T) {
	l1 := []string{"a", "b"}
	i1 := CopySlice(l1)
	assert.Len(t, i1, 2)

	assert.Nil(t, CopySlice(nil))

	var l2 []string
	assert.Nil(t, CopySlice(l2))
}

func TestIsSlice(t *testing.T) {
	assert.True(t, IsSlice([]string{}))
	assert.False(t, IsSlice(5))
}

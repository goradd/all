package any

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	assert.Equal(t, "yes", If(true, "yes", "no"))
	assert.Equal(t, 1, If(false, 2, 1))
	assert.Equal(t, 1, If("", 2, 1))
	assert.Equal(t, 2, If("a", 2, 1))
	assert.Equal(t, "a", If("a", "a", "b"))
}

func ExampleIf() {
	i := 1
	s := ""

	fmt.Println(If(i == 1, 1, 2))
	fmt.Println(If(s, "yes", "no"))

	// Output: 1
	// no
}

func ExampleOr() {
	fmt.Println(Or("yes", "no"))
	fmt.Println(Or("", "no"))

	// Output: yes
	// no
}

func TestIsNil(t *testing.T) {
	var m map[int]int
	assert.True(t, IsNil(nil))
	assert.True(t, IsNil(m))
	assert.False(t, IsNil(5))

	var i any
	assert.True(t, IsNil(i))

	var c chan int
	assert.True(t, IsNil(c))

	var f func(a int)
	assert.True(t, IsNil(f))
}

func TestZero(t *testing.T) {
	assert.Equal(t, "", Zero[string]())
	assert.Equal(t, 0, Zero[int]())
}

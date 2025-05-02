package all

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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
	assert.True(t, Zero[time.Time]().IsZero())
}

func TestIsInteger(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"int", int(1), true},
		{"int8", int8(1), true},
		{"int16", int16(1), true},
		{"int32", int32(1), true},
		{"int64", int64(1), true},
		{"uint", uint(1), true},
		{"uint8", uint8(1), true},
		{"uint16", uint16(1), true},
		{"uint32", uint32(1), true},
		{"uint64", uint64(1), true},
		{"string", "hello", false},
		{"float64", float64(1.0), false},
		{"bool", true, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsInteger(tt.input)
			if got != tt.want {
				t.Errorf("IsInteger(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		{"float32", float32(1.5), true},
		{"float64", float64(1.5), true},
		{"int", int(2), false},
		{"string", "test", false},
		{"bool", false, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsFloat(tt.input)
			if got != tt.want {
				t.Errorf("IsFloat(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

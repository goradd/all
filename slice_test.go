package anyutil

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

type MyType int

func TestMapSlice(t *testing.T) {
	l1 := []string{"a", "b"}
	var i1 []any
	var nilSlice []any
	i1 = MapSlice[any](l1)
	assert.Len(t, i1, 2)
	assert.Equal(t, "a", i1[0].(string))

	assert.Nil(t, MapSlice[any](nilSlice))

	var l2 []string
	assert.Nil(t, MapSlice[any](l2))

	i3 := []any{1, 2}
	i4 := MapSlice[int](i3)
	assert.True(t, i4[0] == 1)

	m := []MyType{1, 2}
	m2 := MapSlice[int](m)
	assert.True(t, m2[0] == 1)
}

func TestIsSlice(t *testing.T) {
	assert.True(t, IsSlice([]string{}))
	assert.False(t, IsSlice(5))
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name     string
		values   []any
		sep      string
		expected string
	}{
		{"Empty slice", []any{}, ", ", ""},
		{"Single element", []any{"Hello"}, ", ", "Hello"},
		{"Multiple strings", []any{"Hello", "World"}, " ", "Hello World"},
		{"Multiple numbers", []any{1, 2, 3}, "-", "1-2-3"},
		{"Mixed types", []any{"Go", 2024, true}, " | ", "Go | 2024 | true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.values, tt.sep)
			if result != tt.expected {
				t.Errorf("Join(%v, %q) = %q; want %q", tt.values, tt.sep, result, tt.expected)
			}
		})
	}
}

func TestMapSliceFunc(t *testing.T) {
	t.Run("int to string", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		result := MapSliceFunc(input, func(i int) string {
			return strconv.Itoa(i)
		})
		if len(result) != len(expected) {
			t.Fatalf("expected length %d, got %d", len(expected), len(result))
		}
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("at index %d: expected %q, got %q", i, expected[i], result[i])
			}
		}
	})

	t.Run("string to upper-case", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := []string{"A", "B", "C"}
		result := MapSliceFunc(input, strings.ToUpper)
		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("at index %d: expected %q, got %q", i, expected[i], result[i])
			}
		}
	})

	t.Run("empty input", func(t *testing.T) {
		input := []float64{}
		result := MapSliceFunc(input, func(f float64) string {
			return strconv.FormatFloat(f, 'f', -1, 64)
		})
		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})
}

func BenchmarkMapSliceReflect(b *testing.B) {
	// Setup: slice of alias type
	input := make([]MyType, 1000)
	for i := 0; i < len(input); i++ {
		input[i] = MyType(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MapSlice[int](input)
	}
}

func BenchmarkManualCast(b *testing.B) {
	// Setup: same slice
	input := make([]MyType, 1000)
	for i := 0; i < len(input); i++ {
		input[i] = MyType(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		output := make([]int, len(input))
		for j, v := range input {
			output[j] = int(v)
		}
	}
}

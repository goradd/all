package any

import (
	"reflect"
	"testing"
)

func TestStringMap(t *testing.T) {
	tests := []struct {
		name  string
		m     map[string]any
		wantO map[string]string
	}{
		{"strings", map[string]any{"a": "A", "b": "B"}, map[string]string{"a": "A", "b": "B"}},
		{"ints", map[string]any{"a": 1, "b": 2}, map[string]string{"a": "1", "b": "2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO := StringMap(tt.m); !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("AnyMapToStrings() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		m     map[string]string
		wantO map[string]any
	}{
		{"basic", map[string]string{"a": "A", "b": "B"}, map[string]any{"a": "A", "b": "B"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO := Map(tt.m); !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("MapToAny() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

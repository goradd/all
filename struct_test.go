package any

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldMap(t *testing.T) {
	a := struct {
		A int
		B string
		C float32
		d float64
	}{
		1, "a", 3.4, 6.7,
	}
	b := FieldMap(a)
	assert.Equal(t, 1, b["A"])
	assert.Equal(t, "a", b["B"])
	assert.Empty(t, b["d"])
}

func TestSetFields(t *testing.T) {
	a := struct {
		A int
		B string
		C float32
		d float64
	}{
		1, "a", 3.4, 6.7,
	}
	b := FieldMap(a)
	c := a
	c.A = 2
	e := SetFields(&c, b)
	assert.NoError(t, e)
	assert.Equal(t, 1, c.A)
	assert.Empty(t, b["d"])

	p := FieldMap(&a)
	assert.Equal(t, "a", p["B"])

	assert.Error(t, SetFields("", b))
}

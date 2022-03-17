package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsErr(t *testing.T) {
	p := &Product{Price: 1.22, SKU: "abc-bcd-cde"}

	v := NewValidation()
	errs := v.Validate(p)
	assert.Len(t, errs, 1)
}

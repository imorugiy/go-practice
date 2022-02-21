package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{Name: "Tea", Price: 1.00, SKU: "abc-abd-ddd"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

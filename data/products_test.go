package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Espresso",
		Price: 10.2,
		SKU:   "Esp-caff-strong",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

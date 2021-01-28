package main

import (
	"fmt"
	"testing"

	"github.com/DeconvFFT/goMicroservicesbasics/sdk/client"
	"github.com/DeconvFFT/goMicroservicesbasics/sdk/client/products"
)

func TestClient(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prod)

}

package main

import (
	"fmt"
	"go-practice/sdk/client"
	"go-practice/sdk/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	config := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, config)
	params := products.NewListProductsParams()
	products, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", products.GetPayload()[0])
	t.Fail()
}

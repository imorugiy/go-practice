package handlers

import (
	"go-practice/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
// 		200: productResponse
//  	422: errorValidation
//  	501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	data.AddProduct(product)
}

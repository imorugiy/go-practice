package handlers

import (
	"go-practice/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a product
//
// responses:
// 		201: noContentResponse
// 		404: errorResponse
// 		422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("Handle PUT Product", product.ID)

	err := data.UpdateProduct(product)
	if err == data.ErrProductNotFound {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found"}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handlers

import (
	"go-practice/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//  	201: noContentResponse
// 		404: errorResponse
// 		501: errorResponse

// DeleteProduct deletes a product from a data store
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

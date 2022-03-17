package handlers

import (
	"go-practice/data"
	"net/http"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 		200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	w.Header().Add("Content-Type", "application/json")

	lp := data.GetProducts()

	err := data.ToJSON(lp, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a single product
// responses:
// 		200: productResponse
//  	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("Handle GET Single Product")

	prod, err := data.GetProductByID(id)

	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

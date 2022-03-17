// Package classification Products API
//
// Documentation for Products API
//
// 		Schemes: http
// 		BasePath: /
// 		Version: 0.0.1
//
// 		Consumes:
// 		- application/json
//
// 		Produces:
// 		- application/json
//
// swagger:meta
package handlers

import "go-practice/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product to Update or Create
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct listSingleProduct
type productIDParameterWrapper struct {
	// The id of the product for which operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

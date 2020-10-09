package handlers

import (
	"net/http"

	"github.com/DeconvFFT/goMicroservicesbasics/data"
)

//  swagger:route GET /products products listProducts
//  Returns a list of products
//  responses :
//  200: productsResponse

//  GetProducts returns the products from datastore
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	//fetch products from data store
	lp := data.GetProducts()

	//serialise list to JSON
	err := data.ToJSON(lp, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

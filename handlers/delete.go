package handlers

import (
	"net/http"
	"strconv"

	"github.com/DeconvFFT/goMicroservicesbasics/data"
	"github.com/gorilla/mux"
)

//  swagger:route DELETE /products/{id} products Delete
//  Deletes product from datastore
//  responses:
//    201: noContentResponse

//  Delete deletes a product from datastore
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle DELETE products", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/DeconvFFT/goMicroservicesbasics/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// by default go calls handler's serveHTTP method

	// to handle get method
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// add a new product
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	//handle update

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		reqEx := regexp.MustCompile(`/([0-9]+)`)
		g := reqEx.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URL more than one id")

			http.Error(rw, "Invalid URL more than one id", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URL more than one capture group")

			http.Error(rw, "Invalid URL more than one capture group", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw, "Invalid URL unable to convert to number", http.StatusBadRequest)
			return
		}

		p.l.Println("got id", id)
		p.UpdateProducts(id, rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

//getProducts returns products from datastore
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// get data from datastore
	lp := data.GetProducts()

	//serialise the list into json
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to Marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

func (p *Products) UpdateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

package handlers

import (
	"log"
	"net/http"

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

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to Marshal json", http.StatusInternalServerError)
	}
}

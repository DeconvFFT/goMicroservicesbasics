package handlers

import (
	"encoding/json"
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
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to Marshal json", http.StatusInternalServerError)
	}
	rw.Write(d)
}

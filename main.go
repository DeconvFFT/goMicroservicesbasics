package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "goodbye")

	})
	http.ListenAndServe(":8080", nil)
}

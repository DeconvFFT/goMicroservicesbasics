package data

import "time"

// use `json : fieldname omitempty (only if you want to ignore if empty)` or `json "-" (to completely ignore a field)`

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price`
	SKU         string  `json:sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy Milk Coffee",
		Price:       2.45,
		SKU:         "latte123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          1,
		Name:        "Espresso ",
		Description: "Short and string coffee without milk",
		Price:       2.21,
		SKU:         "espresso123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() []*Product {
	return productList
}

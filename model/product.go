package model

import "github.com/satori/go.uuid"

type Product struct {
	ID string
	Name string
}

func NewProduct() *Product {
	product := Product{}
	product.ID = uuid.NewV4().String()
	return &product
}

type Products struct {
	Product []Product
}

func (p *Products) Add (product *Product) {
	p.Product = append(p.Product, *product)
}

func NewProducts() *Products {
	return &Products{}
}
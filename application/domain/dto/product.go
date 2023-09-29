package dto

import "github.com/Guilherme-Joviniano/go-hexagonal/application/domain"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *domain.Product) (*domain.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()
	if err != nil {
		return &domain.Product{}, err
	}
	return product, nil
}

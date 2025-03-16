package products

import "time"

type Product struct {
	ProductName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// New is example of factory pattern
func (p *Product) New() *Product {
	return &Product{
		ProductName: p.ProductName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

package models

import (
	"fmt"
	"sync"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
	mu    sync.Mutex
}

func NewProduct(id int, name string, price float64, stock int) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

func (p *Product) DecreaseStock() bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.Stock > 0 {
		p.Stock--

		return true
	}
	return false
}

func (p *Product) IncreaseStock() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Stock++
}

func (p *Product) GetFormattedPrice() string {
	return fmt.Sprintf("Rp. %.0f", p.Price)
}

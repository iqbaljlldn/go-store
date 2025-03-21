package models

import "sync"

type Cashier struct {
	ID               int
	Name             string
	IsAvailable      bool
	TransactionCount int
	TotalSales       float64
	mu               sync.Mutex
}

func NewCashier(id int, name string) *Cashier {
	return &Cashier{
		ID:          id,
		Name:        name,
		IsAvailable: true,
	}
}

func (c *Cashier) SetAvailability(isAvailable bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.IsAvailable = isAvailable
}

func (c *Cashier) IsAvailableForTransaction() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.IsAvailable
}

func (c *Cashier) RecordTransaction(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.TransactionCount++
	c.TotalSales += amount
}

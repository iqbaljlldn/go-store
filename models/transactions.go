package models

import "time"

type Transaction struct {
	ID          int
	Customer    *Customer
	Cashier     *Cashier
	TotalAmount float64
	Discount    float64
	FinalAmount float64
	Items       []*CartItem
	Timestamps  time.Time
	Status      string
}

func NewTransaction(id int, customer *Customer, cashier *Cashier) *Transaction {
	return &Transaction{
		ID:         id,
		Customer:   customer,
		Cashier:    cashier,
		Items:      customer.ShoppingCart,
		Timestamps: time.Now(),
	}
}

func (t *Transaction) IsSuccessfull() bool {
	return t.Status == "Success" || t.Status == "Success (Retry)"
}

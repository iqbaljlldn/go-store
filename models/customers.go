package models

import "time"

type Customer struct {
	ID           int
	Name         string
	ShoppingCart []*CartItem
	ProcessTime  time.Duration
}

func NewCustomer(id int, name string, processTime time.Duration) *Customer {
	return &Customer{
		ID:           id,
		Name:         name,
		ShoppingCart: []*CartItem{},
		ProcessTime:  processTime,
	}
}

func (c *Customer) AddToCart(product *Product, quantity int) {
	cartItem := &CartItem{
		Product:  product,
		Quantity: quantity,
	}
	c.ShoppingCart = append(c.ShoppingCart, cartItem)
}

func (c *Customer) CartItemCount() int {
	var total int
	for _, item := range c.ShoppingCart {
		total += item.Quantity
	}

	return total
}

func (c *Customer) CartTotalPrice() float64 {
	var total float64
	for _, item := range c.ShoppingCart {
		total += item.Product.Price * float64(item.Quantity)
	}

	return total
}

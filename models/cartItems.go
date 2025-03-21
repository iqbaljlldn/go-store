package models

type CartItem struct {
	Product  *Product
	Quantity int
}

func NewCartItem(product *Product, quantity int) *CartItem {
	return &CartItem{
		Product:  product,
		Quantity: quantity,
	}
}

func (c *CartItem) TotalPrice() float64 {
	return c.Product.Price * float64(c.Quantity)
}

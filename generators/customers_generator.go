package generators

import (
	"class/task/models"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type CustomerGenerator interface {
	Generate(count int, products []*models.Product) []*models.Customer
}

type RandomCustomerGenerator struct{}

func NewRandomCustomerGenerator() *RandomCustomerGenerator {
	return &RandomCustomerGenerator{}
}

func (g *RandomCustomerGenerator) Generate(count int, products []*models.Product) []*models.Customer {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // PRNG yang lebih aman
	customers := make([]*models.Customer, 0, count)        // Menghindari elemen nil

	for i := 1; i <= count; i++ {
		customer := models.NewCustomer(
			i,
			gofakeit.Name(),
			time.Duration(rng.Intn(4000)+1000)*time.Millisecond,
		)

		numItems := rng.Intn(5) + 2
		selectedProducts := make(map[int]struct{}) // Lebih hemat memori

		availableProducts := make([]*models.Product, 0, len(products))
		for _, p := range products {
			if p.Stock > 0 {
				availableProducts = append(availableProducts, p)
			}
		}

		// Jika tidak ada produk yang tersedia, lewati customer ini
		if len(availableProducts) == 0 {
			continue
		}

		// Pilih produk unik dengan stok tersedia
		for j := 0; j < numItems && len(availableProducts) > 0; j++ {
			productIndex := rng.Intn(len(availableProducts))
			product := availableProducts[productIndex]

			if _, exists := selectedProducts[product.ID]; exists {
				continue
			}

			quantity := rng.Intn(2) + 1
			if quantity > product.Stock {
				quantity = product.Stock
			}

			customer.AddToCart(product, quantity)
			selectedProducts[product.ID] = struct{}{}

			// Jika stok habis, hapus produk dari daftar yang bisa dipilih
			if product.Stock <= 0 {
				availableProducts = append(availableProducts[:productIndex], availableProducts[productIndex+1:]...)
			}
		}

		// Hanya tambahkan customer jika dia memiliki produk di keranjang
		if len(customer.ShoppingCart) > 0 {
			customers = append(customers, customer)
		}
	}

	return customers
}

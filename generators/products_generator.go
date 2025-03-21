package generators

import (
	"class/task/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type ProductGenerator interface {
	Generate(count int) []*models.Product
}

type RandomProductGenerator struct{}

func NewRandomProductGenerator() *RandomProductGenerator {
	return &RandomProductGenerator{}
}

func (g *RandomProductGenerator) Generate(count int) []*models.Product {
	// Inisialisasi PRNG yang lebih efisien
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	faker := gofakeit.NewCrypto()

	categories := []string{"Electronics", "Food", "Clothing", "Groceries", "Home", "Toys"}
	products := make([]*models.Product, 0, count) // Efisiensi alokasi slice

	for i := 1; i <= count; i++ {
		category := categories[rng.Intn(len(categories))]

		product := models.NewProduct(
			i,
			fmt.Sprintf("%s - %s", category, faker.ProductName()),
			float64(rng.Intn(90000)+10000)/10,
			rng.Intn(90)+10,
		)

		products = append(products, product)
	}

	return products
}

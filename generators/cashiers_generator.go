package generators

import (
	"class/task/models"

	"github.com/brianvoe/gofakeit/v6"
)

type CashierGenerator interface {
	Generate(count int) []*models.Cashier
}

type RandomCashierGenerator struct{}

func NewRandomCashierGenerator() *RandomCashierGenerator {
	return &RandomCashierGenerator{}
}

func (g *RandomCashierGenerator) Generate(count int) []*models.Cashier {
	if count > 5 {
		count = 5
	}

	cashiers := make([]*models.Cashier, count)
	for i := 0; i < count; i++ {
		cashiers[i] = models.NewCashier(
			i+1,
			gofakeit.Name(),
		)
	}

	return cashiers
}

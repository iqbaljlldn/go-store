package factory

import (
	"class/task/generators"
	"class/task/models"
	"class/task/services"
)

type StoreFactory struct {
	ProductGenerator   generators.ProductGenerator
	CashierGenerator   generators.CashierGenerator
	CustomerGenerator  generators.CustomerGenerator
	TransactionService *services.TransactionService
	ReceiptService     services.ReceiptService
}

func NewStoreFactory(
	productGenerator generators.ProductGenerator,
	cashierGenerator generators.CashierGenerator,
	customerGenerator generators.CustomerGenerator,
	transactionService *services.TransactionService,
	receiptService services.ReceiptService,
) *StoreFactory {
	return &StoreFactory{
		ProductGenerator:   productGenerator,
		CashierGenerator:   cashierGenerator,
		CustomerGenerator:  customerGenerator,
		TransactionService: transactionService,
		ReceiptService:     receiptService,
	}
}

func (f *StoreFactory) CreateStore(name string, numCashiers int, numCustomers int, numProducts int, config *models.StoreConfig) *models.Store {
	store := models.NewStore(name, config)

	products := f.ProductGenerator.Generate(numProducts)
	store.SetProducts(products)

	cashiers := f.CashierGenerator.Generate(numCashiers)
	store.SetCashiers(cashiers)

	customers := f.CustomerGenerator.Generate(numCustomers, products)
	store.SetCustomers(customers)

	return store
}

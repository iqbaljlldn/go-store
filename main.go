package main

import (
	"class/task/factory"
	"class/task/generators"
	"class/task/models"
	"class/task/services"
	"time"
)

func main() {
	storeName := "SRC"
	numCashiers := 5
	numCustomers := 100
	numProducts := 100000

	config := models.NewStoreConfig(
		3*time.Second,
		3,
		10.0,
	)

	productGenerator := generators.NewRandomProductGenerator()
	cashierGenerator := generators.NewRandomCashierGenerator()
	customerGenerator := generators.NewRandomCustomerGenerator()

	transactionService := services.NewTransactionService()
	receiptService := services.NewReceiptService()

	storeFactory := factory.NewStoreFactory(
		productGenerator,
		cashierGenerator,
		customerGenerator,
		transactionService,
		receiptService,
	)

	store := storeFactory.CreateStore(storeName, numCashiers, numCustomers, numProducts, config)
	simulation := NewStoreSimulation(store)
	simulation.Run()

	summaryPrinter := NewStoreSummaryPrinter(store)
	summaryPrinter.PrintSummary()
}

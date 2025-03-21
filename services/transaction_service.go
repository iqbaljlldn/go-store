package services

import (
	"class/task/models"
	"context"
	"time"
)

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) CalculateTotal(customer *models.Customer, config *models.StoreConfig) (float64, float64, float64) {
	var totalAmount float64
	var totalItems int

	for _, item := range customer.ShoppingCart {
		totalAmount += item.TotalPrice()
		totalItems += item.Quantity
	}

	var discount float64
	if totalItems >= config.DiscountThreshold {
		discount = totalAmount * config.DiscountPercent / 100
	}

	finalAmount := totalAmount - discount

	return finalAmount, discount, totalAmount
}

func (t *TransactionService) ProcessTransaction(store *models.Store, cashier *models.Cashier, customer *models.Customer, receiptService ReceiptService) *models.Transaction {
	transactionID := store.GetNextTransactionID()
	transaction := models.NewTransaction(transactionID, customer, cashier)

	ctx, cancel := context.WithTimeout(context.Background(), store.Config.TransactionTimeout)
	defer cancel()

	processChan := make(chan bool)

	go func() {
		time.Sleep(customer.ProcessTime)

		allItemsAvailable := true
		for _, item := range customer.ShoppingCart {
			for i := 0; i < item.Quantity; i++ {
				if !item.Product.DecreaseStock() {
					allItemsAvailable = false
					break
				}
			}

			if !allItemsAvailable {
				break
			}
		}

		if allItemsAvailable {
			totalAmount, discount, finalAmount := t.CalculateTotal(customer, store.Config)
			transaction.TotalAmount = totalAmount
			transaction.Discount = discount
			transaction.FinalAmount = finalAmount

			processChan <- true
		} else {
			processChan <- false
		}
	}()

	select {
	case success := <-processChan:
		if success {
			transaction.Status = "Success"

			cashier.RecordTransaction(transaction.FinalAmount)

			receiptService.PrintReceipt(store, transaction)
		} else {
			transaction.Status = "Failed"

			store.QueueFailedTransaction(transaction)
		}
	case <-ctx.Done():
		transaction.Status = "Timeout"
		store.QueueFailedTransaction(transaction)
	}

	store.AddTransactions(transaction)

	return transaction
}

func (t *TransactionService) RetryTransaction(store *models.Store, failedTransaction *models.Transaction, cashier *models.Cashier, receiptService ReceiptService) {
	transactionID := store.GetNextTransactionID()

	newTransaction := models.NewTransaction(transactionID, failedTransaction.Customer, cashier)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		defer cancel()

		// Use a timer channel to simulate processing time
		processDone := make(chan bool)
		go func() {
			time.Sleep(failedTransaction.Customer.ProcessTime / 2)
			processDone <- true
		}()

		// Wait for either processing to finish or context timeout
		select {
		case <-processDone:
			// Processing completed successfully
			totalAmount, discount, finalAmount := t.CalculateTotal(failedTransaction.Customer, store.Config)
			newTransaction.TotalAmount = totalAmount
			newTransaction.Discount = discount
			newTransaction.FinalAmount = finalAmount
			newTransaction.Status = "Success (Retry)"
			cashier.RecordTransaction(newTransaction.FinalAmount)
			receiptService.PrintReceipt(store, newTransaction)
			store.AddTransactions(newTransaction)
		case <-ctx.Done():
			// Timeout occurred
			newTransaction.Status = "Timeout (Retry)"
			store.AddTransactions(newTransaction)
		}

		// Make cashier available again in either case
		cashier.SetAvailability(true)
	}()
}

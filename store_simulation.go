package main

import (
	"class/task/models"
	"class/task/services"
	"fmt"
	"sync"
	"time"
)

type StoreSimulation struct {
	Store              *models.Store
	TransactionService *services.TransactionService
	ReceiptService     services.ReceiptService
}

func NewStoreSimulation(store *models.Store) *StoreSimulation {
	return &StoreSimulation{
		Store:              store,
		TransactionService: services.NewTransactionService(),
		ReceiptService:     services.NewReceiptService(),
	}
}

func (s *StoreSimulation) Run() {
	fmt.Printf("Starting %s Store simulation with %d cashiers and %d customers\n", s.Store.Name, len(s.Store.Cashiers), len(s.Store.Customers))

	go s.handleFailedTransactions()

	customerQueue := make(chan *models.Customer, len(s.Store.Customers))
	for _, customer := range s.Store.Customers {
		customerQueue <- customer
	}
	close(customerQueue)

	var wg sync.WaitGroup

	for _, cashier := range s.Store.Cashiers {
		wg.Add(1)
		go func(c *models.Cashier) {
			defer wg.Done()

			for customer := range customerQueue {
				c.SetAvailability(false)
				fmt.Printf("Cashier %s is serving customer %s\n", c.Name, customer.Name)
				s.TransactionService.ProcessTransaction(s.Store, c, customer, s.ReceiptService)
				c.SetAvailability(true)
			}
		}(cashier)
	}
	wg.Wait()

	time.Sleep(2 * time.Second)

	s.Store.CloseFailedTransactionQueue()
}

func (s *StoreSimulation) handleFailedTransactions() {
	for transaction := range s.Store.FailedTransaction {
		var cashier *models.Cashier
		for {
			for _, c := range s.Store.Cashiers {
				if c.IsAvailableForTransaction() {
					cashier = c
					c.SetAvailability(false)
					break
				}
			}

			if cashier != nil {
				break
			}

			time.Sleep(500 * time.Millisecond)
		}

		fmt.Printf("Retrying failed transaction #%d for customers %s with cashier %s\n", transaction.ID, transaction.Customer.Name, transaction.Cashier.Name)

		s.TransactionService.RetryTransaction(s.Store, transaction, cashier, s.ReceiptService)
	}
}

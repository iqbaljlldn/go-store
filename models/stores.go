package models

import (
	"sync"
	"time"
)

type StoreConfig struct {
	TransactionTimeout time.Duration
	DiscountThreshold  int
	DiscountPercent    float64
}

func NewStoreConfig(transactionTimeout time.Duration, discountThreshold int, discountPercent float64) *StoreConfig {
	return &StoreConfig{
		TransactionTimeout: transactionTimeout,
		DiscountThreshold:  discountThreshold,
		DiscountPercent:    discountPercent,
	}
}

type Store struct {
	Name               string
	Products           []*Product
	Cashiers           []*Cashier
	Customers          []*Customer
	Transactions       []*Transaction
	Config             *StoreConfig
	FailedTransaction  chan *Transaction
	TransactionCounter int
	mu                 sync.Mutex
}

func NewStore(name string, config *StoreConfig) *Store {
	return &Store{
		Name:              name,
		Config:            config,
		FailedTransaction: make(chan *Transaction),
		mu:                sync.Mutex{},
	}
}

func (s *Store) SetProducts(products []*Product) {
	s.Products = products
}

func (s *Store) SetCashiers(cashier []*Cashier) {
	s.Cashiers = cashier
}

func (s *Store) SetCustomers(customer []*Customer) {
	s.Customers = customer
}

func (s *Store) AddTransactions(transaction *Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Transactions = append(s.Transactions, transaction)
}

func (s *Store) GetNextTransactionID() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TransactionCounter++

	return s.TransactionCounter
}

func (s *Store) QueueFailedTransaction(transaction *Transaction) {
	s.FailedTransaction <- transaction
}

func (s *Store) CloseFailedTransactionQueue() {
	close(s.FailedTransaction)
}

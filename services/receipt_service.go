package services

import (
	"class/task/models"
	"fmt"
)

type ReceiptService interface {
	PrintReceipt(store *models.Store, transaction *models.Transaction)
}

type DefaultReceiptService struct{}

func NewReceiptService() *DefaultReceiptService {
	return &DefaultReceiptService{}
}

func (s *DefaultReceiptService) PrintReceipt(store *models.Store, t *models.Transaction) {
	fmt.Printf("\n=== Receipt ===\n")
	fmt.Printf("Store: %s\n", store.Name)
	fmt.Printf("Transaction #%d\n", t.ID)
	fmt.Printf("Date: %s\n", t.Timestamps.Format("2006-01-02 15:04:05"))
	fmt.Printf("Cashier: %s\n", t.Cashier.Name)
	fmt.Printf("Customer: %s\n", t.Customer.Name)
	fmt.Printf("--------------------\n")
	fmt.Printf("Items:\n")

	for _, item := range t.Items {
		fmt.Printf("	%s x%d: Rp. %.0f\n", item.Product.Name, item.Quantity, item.TotalPrice())
	}

	fmt.Printf("---------------------\n")
	fmt.Printf("Subtotal: Rp.%f\n", t.TotalAmount)
	if t.Discount > 0 {
		fmt.Printf("Discount (%.0f%%): -Rp. %.0f\n", store.Config.DiscountPercent, t.Discount)
	}

	fmt.Printf("Total: Rp. %0.f\n", t.FinalAmount)
	fmt.Printf("Status: %s\n", t.Status)
	fmt.Printf("=================\n")
}

package main

import (
	"class/task/models"
	"fmt"
)

type StoreSummaryPrinter struct {
	Store *models.Store
}

func NewStoreSummaryPrinter(store *models.Store) *StoreSummaryPrinter {
	return &StoreSummaryPrinter{Store: store}
}

func (p *StoreSummaryPrinter) PrintSummary() {
	var successCount, failedCount, timeoutCount int
	var totalSales float64

	for _, t := range p.Store.Transactions {
		if t.IsSuccessfull() {
			successCount++
			totalSales += t.FinalAmount
		} else if t.Status == "Failed" {
			failedCount++
		} else if t.Status == "Timeout" {
			timeoutCount++
		}
	}

	fmt.Printf("\n======= SIMULATION SUMMARY =======\n")
	fmt.Printf("Store: %s\n", p.Store.Name)
	fmt.Printf("Total Transactions: %d\n", len(p.Store.Transactions))
	fmt.Printf("Successful Transactions: %d\n", successCount)
	fmt.Printf("Failed Transactions: %d\n", failedCount)
	fmt.Printf("Timeout Transactions: %d\n", timeoutCount)
	fmt.Printf("Total sales: Rp. %.0f\n\n", totalSales)

	for _, cashier := range p.Store.Cashiers {
		fmt.Printf("Cashier %s (ID: %d):\n", cashier.Name, cashier.ID)
		fmt.Printf(" - Transactions: %d\n", cashier.TransactionCount)
		fmt.Printf(" - Total Sales: Rp.%.0f\n\n", cashier.TotalSales)
	}

	fmt.Printf("===================================\n")
}

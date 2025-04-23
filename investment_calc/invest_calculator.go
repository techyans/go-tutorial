package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax: ")
	fmt.Scan(&tax)

	profit := revenue - expenses
	netProfit := profit - (profit * tax / 100)
	tax_amt := profit - netProfit
	profilt_raion := profit / netProfit

	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Net Profit: %.2f\n", netProfit)
	fmt.Printf("Tax Amount: %.2f \n", tax_amt)
	fmt.Printf("Profit Ratio: %.2f", profilt_raion)
}

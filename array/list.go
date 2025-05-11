package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	price := []float64{99.2, 23, 24.1}

	price = append(price, 100.2, 25.3)
	price = append([]float64{100}, price...)

	fmt.Println(price)
}

/*func main() {
	var ProductName [4]string = [4]string{"A", "B", "C", "D"}
	prices := [4]float64{1.2, 2.3, 3.4, 4.5}

	fmt.Println(prices[2])
	ProductName[2] = "EAt"
	fmt.Println(ProductName)

	FinalPrice := prices[1:4]
	featuredPrice := prices[1:]
	featuredPrice[0] = 199.99
	fmt.Println(FinalPrice)
	fmt.Println(featuredPrice)
	fmt.Println(prices)

	fmt.Println(len(featuredPrice), cap(featuredPrice))
}
*/

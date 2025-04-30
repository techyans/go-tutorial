package main

import "fmt"

func main() {
	age := 10
	adultAge := &age
	fmt.Println(*adultAge)

	AdultAgeFinal := changeAge(adultAge)
	fmt.Println(AdultAgeFinal)
}

func changeAge(age *int) int {
	return *age - 5
}

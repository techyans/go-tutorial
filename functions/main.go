package main

import (
	"fmt"
)

type transfrm func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 50}
	Morenumber := []int{5, 10, 30, 80, 100}
	numbers = DoubleNumber(&numbers, triple)
	fmt.Println(numbers)

	transfn1 := GetTransformfunc(&numbers)
	transfn2 := GetTransformfunc(&Morenumber)

	GetTransfrm1 := DoubleNumber(&numbers, transfn1)
	GetTransfrm2 := DoubleNumber(&Morenumber, transfn2)
	fmt.Println(GetTransfrm1, GetTransfrm2)
}

func DoubleNumber(numbers *[]int, transform transfrm) []int {
	doubleNum := []int{}

	for _, num := range *numbers {
		doubleNum = append(doubleNum, transform(num))
	}

	return doubleNum
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func GetTransformfunc(number *[]int) transfrm {
	if (*number)[0] == 1 {
		return double
	} else {
		return triple
	}
}

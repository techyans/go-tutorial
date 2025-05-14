package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}
func main() {
	Username := make([]string, 2, 3)

	Username = append(Username, "Arshad")
	Username = append(Username, "Jamal")

	Username[0] = "Arshad Jamal"
	fmt.Println(Username)

	courseRating := make(floatMap, 12)
	courseRating["Go"] = 4.9
	courseRating["Python"] = 4.8

	courseRating.output()
	// fmt.Println(courseRating)

	for index, value := range courseRating {
		fmt.Println("index", index)
		fmt.Println("value", value)
	}

	for key, val := range Username {
		fmt.Println("key", key)
		fmt.Println("val", val)
	}

}

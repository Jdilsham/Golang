package main

import "fmt"

func main() {
	fmt.Println("Welcome to function")
	greeeting()

	//Can not define function inside the another function
	//func greeetingtwo() {
	//	fmt.Println("Hello World")
	//}

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	Total := proAdder(1, 2, 3, 4, 5, 6, 8, 7, 9)
	fmt.Printf("Total is: %v\n", Total)

}

func greeeting() {
	fmt.Println("Hello World")
}

// return type need to ge mentions
func adder(a int, b int) int {
	return a + b
}

//add numbers without knowing how many numbers need to be added

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}

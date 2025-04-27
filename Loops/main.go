package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	//days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	//fmt.Println(days)

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {
	//	fmt.Println("Day", index, ":", day)
	//}

	value := 1
	for value < 10 { //Similar to while loop in other languages

		if value == 5 {
			break
		}

		fmt.Printf("%d ", value)
		value = value + 1
	}
	//Also can use continue

}

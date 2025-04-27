package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	//fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 ad you can open")
	case 2:
		fmt.Println("Dice value is 2 ad you can spot")
	case 3:
		fmt.Println("Dice value is 3 ad you can spot")
	case 4:
		fmt.Println("Dice value is 4 ad you can spot")
	case 5:
		fmt.Println("Dice value is 5 ad you can spot")
	case 6:
		fmt.Println("Dice value is 6 ad you can roll again")
	default:
		fmt.Println("What was not a dice number?")

	}

}

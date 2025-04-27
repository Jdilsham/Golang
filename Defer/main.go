package main

import "fmt"

func main() {
	// First deferred statement - will execute last due to LIFO
	defer fmt.Println("one")
	// Second deferred statement - will execute second-to-last
	defer fmt.Println("two")
	// This prints immediately
	fmt.Println("Welcome to Defer")
	// Defer statements execute after the function returns in Last-In-First-Out order
	// So output will be: "Welcome to Defer" -> numbers from myDefer() -> "two" -> "one"

	// Call myDefer() which demonstrates defer with a loop
	myDefer()
}

func myDefer() {
	// Loop from 0 to 4
	for i := 0; i < 5; i++ {
		// Defers printing of each number
		// Due to LIFO (Last In First Out) nature of defer
		// Numbers will be printed in reverse order: 4,3,2,1,0
		defer fmt.Println(i)
	}
}

package main

import (
	"fmt"
)

func main() {
	// This is an entry point to be exectued.
	// simpleFunction()
	add(23, 20)
}

// This is a *simple-funciton*.
func simpleFunction() {
	fmt.Println("Hello, simple func. here...")
}

// funciton with *parameters*.
// Arguments are specified after the function name, inside the parentheses. You can add as many arguments as you want, just separate them with a comma.
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

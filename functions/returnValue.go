package main

import (
	"errors"
	"fmt"
)

func main() {
	// result := square(3)
	// fmt.Println(result)

	// a, b := swap(2, 5)
	// fmt.Println(a, b)

	// l, m := split(27)
	// fmt.Println(l, m)

	res, err := divide(10, 2)
	// fmt.Println(res, err) // It will print the both.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

// Now this is single return values in funciton.
func square(x int) int {
	return x * x
}

// multiple return values in functions.
func swap(x int, y int) (int, int) {
	return y, x
}

// named return values.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Returning an errors.
// It's a common practice in Go to return an error value as the last return value of a function, which allows you to handle errors effectively.
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

// Blank identifiers.
// If you want to ignore one of the returned values from a function, you can use the blank identifier _.

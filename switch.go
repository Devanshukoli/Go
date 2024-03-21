package main

import (
	"fmt"
	"time"
)

func main() {
	// switch_statement()
	// switch_conditional()
	err_switch()
}

func switch_statement() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean the damn house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Checkup with doctor.")
	case 20, 21, 22, 23, 24, 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Tonight is party. boys.")
	default:
		fmt.Println("I don't give a fk.")
	}
}

func switch_conditional() {
	today := time.Now()

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func err_switch() {
	//All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error
	a := 3

	switch a {
	case 1:
		fmt.Println("a is one")
	case "b":
		fmt.Println("a is b")
	}
}

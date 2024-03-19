package main

import (
	"fmt"
)

func main() {
	// conditional()
	// if_statement()
	// if_else_statement()
	// else_if_statement()
	nested_if_statement()
}

func conditional() {
	// Greater than
	x := 10
	y := 5
	z := 30
	fmt.Println("x>y", x > y)

	// Not Equal to
	fmt.Println("x!=y", x != y)

	// And operator
	fmt.Println("(x > y) && (y > z)", (x > y) && (y > z))
}

func if_statement() {
	// syntax
	// if (condition) {
	// 	// code to execute if the cond. is true.
	// }
	if 10 > 4 {
		fmt.Println("It's win.")
	}
}

func if_else_statement() {

	age := 21
	if age < 20 {
		fmt.Println("You are underage")
	} else {
		fmt.Println(("you are grown now."))
	}

	// Note : Note that  `} else {}` is the only correct way other wise it will raise an error.
}

func else_if_statement() {
	var num int = 23
	if num > 30 {
		fmt.Println("It's win.")
	} else if num < 25 {
		fmt.Println("it's tie. you know.")
	} else {
		fmt.Println("It's loose.")
	}
}

func nested_if_statement() {
	// You can have `if` statement inside `if` statement that is what called **nested if**.
	weight := 60
	if weight > 50 {
		fmt.Println("Damn, You are weighted.")
		if weight > 100 {
			fmt.Println("You are HEAVY!!!")
		}
	} else {
		fmt.Println("Yup, Fine weight.")
	}
}

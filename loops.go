package main

import (
	"fmt"
)

func main() {
	// loop()
	// while_loop()
	infinite_loop()
	// continue_statment_loop()
	// break_statement_loop()
	// nested_loops()
}

func loop() {
	// simple loop for 1 to 5...
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// loop through 100 nums with 10nums gap.
	for i := 0; i < 100; i += 10 {
		fmt.Println(i)
	}
}

func while_loop() {
	// if you skip the initialization and counter value then you get while loop.
	i := 1
	for i < 5 {
		i *= 2
		fmt.Println(i)
	}
}

func infinite_loop() {
	// if I don't add any statements here then it will go to infinite loop.
	sum := 0
	for {
		sum++
		fmt.Println("inside loop..", sum)
	}
	// It will never reach here.
	fmt.Println("outside loop...", sum)
}

func continue_statment_loop() {
	//The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
	for i := 0; i < 5; i++ {
		if i == 3 {
			print("horray, it's 3\n")
			continue
		}
		fmt.Println(i)
	}
}

func break_statement_loop() {
	// This example breaks out of the loop when i is equal to 3:
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

func nested_loops() {
	first := [2]string{"big", "tasty"}
	last := [3]string{"apple", "banana", "orange"}

	for i := 0; i < len(first); i++ {
		for j := 0; j < len(last); j++ {
			fmt.Println(first[i], last[j])
		}
	}
}

package main

import (
	"fmt"
)

func main() {
	arithmetic_op()
	assignment_op()
	comparison_op()
	logical_op()
	bitwise_op()
}

func arithmetic_op() {
	// Arithmetic operaitons.
	x := 10
	y := 2
	fmt.Println("Add", x+y)
	fmt.Println("Sub", x-y)
	fmt.Println("Multiple", x*y)
	fmt.Println("Divide", x/y)
	// fmt.Println("Add", x++)
	// fmt.Println("Add", x__)
}

func assignment_op() {
	//Assignment Operators
	x := 10
	x += 3
	fmt.Println("+= ,", x)
	x -= 3
	fmt.Println("-= ,", x)
	x *= 3
	fmt.Println("*= ,", x)
	x /= 3
	fmt.Println("/= ,", x)
	x %= 3
	fmt.Println("%= ,", x)
	x &= 3
	fmt.Println("&= ,", x)
	x |= 3
	fmt.Println("|= ,", x)
	x ^= 3
	fmt.Println("^= ,", x)
	x >>= 3
	fmt.Println(">>= ,", x)
	x <<= 3
	fmt.Println(">>= ,", x)
}

func comparison_op() {
	// Note: The return value of a comparison is either true (1) or false (0).
	var x = 5
	var y = 6
	fmt.Println("x==y", x == y)
	fmt.Println("x!=y", x != y)
	fmt.Println("x>y", x > y)
	fmt.Println("x<y", x < y)
	fmt.Println("x>=y", x >= y)
	fmt.Println("x<=y", x <= y)
}

func logical_op() {
	x := 5
	fmt.Println("x =", x)
	fmt.Println("x< 5 && x > 10", x < 5 && x > 10)
	fmt.Println("x < 5 || x < 4", x < 5 || x < 4)
	fmt.Println("!(x < 5 && x < 10)", !(x < 5 && x < 10))
}

func bitwise_op() {

}

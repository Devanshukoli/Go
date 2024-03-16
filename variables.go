package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 3
	fmt.Println("b-->", b, "c--->", c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	// multiple Variables.
	var p, q, r, s int = 1, 3, 4, 5
	fmt.Println(p, q, r, s)

	var x, y = 6, "Hello"
	n, z := 7, "World!"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(n)
	fmt.Println(z)
}

package main

import "fmt"

func main() {
	var a = 1
	var b = "hello"
	var c = 32
	// 1. Very much *c* like syntax. for priting something.
	fmt.Print(a, " \n")
	fmt.Print(b, "\n")
	fmt.Print(a, " ", b, "\n")
	fmt.Print(a, c) // Here, Print() automatically adds *space* between them if **Neither** are strings.

	// 2. Println() *java* like  syntax. The diff. between Print() and Println() is that it adds space between them and also add new line.
	fmt.Println(a, b)

	// 3. Printf() :
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

}

package main

import (
	"fmt"
)

type Person struct {
	age   int
	name  string
	hobby string
}

func main() {
	fmt.Println(Person{21, "Devanshu", "Learning"})
}

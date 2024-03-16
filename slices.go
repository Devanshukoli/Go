package main

import (
	"fmt"
)

func main() {
	// slice1()
	// slice2()
	slice3()
}

// Creating slice using -> []datatype{value}
func slice1() {
	myslice1 := []int{}
	fmt.Println("Length of empty slice ---->", len(myslice1))
	fmt.Println("Capacity of empty slice ---->", cap(myslice1))
	fmt.Println("empty slice ---->", myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}

// Creating slice using -> make() function
func slice2() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	// Note: If the capacity parameter is not defined, it will be equal to length.
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}

// Creating slice using array.
func slice3() {
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))
}

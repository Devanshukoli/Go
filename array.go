package main

import "fmt"

func main() {
	declaration()
	access_array()
	change_element()
	initialize()
	initial_only_specific_elements()
	length_of_array()
}

// Array declaration of Array.
func declaration() {
	var arr = [3]int{1, 3, 4}
	arr2 := [4]string{"damn", "tea", "cake", "chocalate"}
	arr3 := [...]float64{23.23, 34.2, 23}
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// Access elements of array.
func access_array() {
	var cars = [4]string{"Audi", "BMW", "Buggati", "TATA"}
	fmt.Println(cars)    // all the values should be print.
	fmt.Println(cars[0]) // Audi
}

// Change elements of array.
func change_element() {
	var nums = [3]int{23, 45, 56}
	fmt.Println(nums)
	nums[1] = 83
	fmt.Println(nums)
}

// Array Initialization
func initialize() {
	// If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
	// Tip: The default value for int is 0, and the default value for string is "".
	emotions := [5]string{}
	emotions1 := [5]string{"demo", "split"}
	emotions2 := [5]string{"demo", "split", "known", "all", "black"}
	num1 := [5]int{}              //not initialized
	num2 := [5]int{1, 2}          //partially initialized
	num3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	fmt.Println(emotions)
	fmt.Println(emotions1)
	fmt.Println(emotions2)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
}

// Initialize only specific elements.
func initial_only_specific_elements() {
	star_design := [...]int{1: '*'}
	fmt.Println(star_design)
}

// Length of array.
func length_of_array() {
	var length_check = [2]int{1, 4}
	fmt.Println(len(length_check))
}

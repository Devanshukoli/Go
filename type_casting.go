package main

import (
	"fmt"
	"strconv"
)

func main() {
	// type_casting()
	stringconv()
}

func type_casting() {
	// Type conversion happen when we assign the value of one data type to another.
	var totalsum int = 850
	var number int = 19
	var avg float32

	// avg = float32(totalsum)/float32(number)
	avg = float32(totalsum + number)
	// As you can see here, you have to convert them into same data type otherwise they will not be exectued.
	fmt.Println(avg)

	i := int(23.2334) // Here, it will not let me convert.
	fmt.Println(i)
}

func stringconv() {
	var s string = "kalp"
	v, _ := strconv.Atoi(s)

	fmt.Println(v)
}

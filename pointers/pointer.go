package main

import "fmt"

func main() {
	val := 10
	fmt.Println(val)

	// ptr stores address
	ptr := &val
	fmt.Println(ptr)

	// * used for dereferencing
	fmt.Println(*ptr)

	*ptr = 50 // lets update the value , using pointer
	fmt.Println("After updating val:", val)

	val = 100 // update the value of val
	fmt.Println("After updating pointer:", val)
}

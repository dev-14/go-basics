package main

import (
	"fmt"
)

func calculator(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func main() {
	var a, b int
	// var name string
	fmt.Print("Enter 2 numbers ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("Calculator Function")
	add, sub, mult, div := calculator(a, b)
	fmt.Println("Addition is :- ", add)
	fmt.Println("Substraction is :- ", sub)
	fmt.Println("Multiplication is :- ", mult)
	fmt.Println("Division is :- ", div)

}

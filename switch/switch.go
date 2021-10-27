package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number between 1 and 4")
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println("case one")
	case 2:
		fmt.Println("case two")
	case 3:
		fmt.Println("case three")
	case 4:
		fmt.Println("case four")
	default:
		fmt.Println("number out of range")
	}

}

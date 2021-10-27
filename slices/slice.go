package main

import "fmt"

func main() {
	s := make([]int, 10)
	fmt.Println("s initialized as ", s)
	s[1] = 1
	s[2] = 2
	s[3] = 3
	s[4] = 4
	s[5] = 5
	s[6] = 6
	sLength := len(s)
	fmt.Println(s)
	fmt.Println("length of slice is: ", sLength)
}

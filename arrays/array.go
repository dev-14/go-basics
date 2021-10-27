package main

import "fmt"

func main() {
	fmt.Println("1 dimensional array")
	var intArr [4]int
	fmt.Println(intArr)

	fmt.Println("2 dimensional array")
	var doubleArray [3][3]int
	doubleArray[1][1] = 5
	doubleArray[2][2] = 6
	fmt.Println(doubleArray)
}

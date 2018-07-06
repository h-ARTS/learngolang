package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	mySlice = append(mySlice, 8)
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
	fmt.Println(mySlice)
}

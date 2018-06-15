package main

import (
	"fmt"
)

func main() {

	mySlice := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2]) // 83 ascii
}

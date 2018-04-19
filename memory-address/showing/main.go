package main

import (
	"fmt"
)

func main() {

	a := 43

	fmt.Println("a's memory address: ", &a) // Ampersand + variable
	fmt.Printf("In Decimal: %d \n", &a)     // Ampersand + variable
}

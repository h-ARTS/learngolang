package main

import (
	"fmt"
)

func zero(z *int) {
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	one(&x)
	fmt.Println(x) // x is now 0, because of passing a value on the same memory address!
}

func one(c *int) {
	fmt.Printf("Here: %v", c)
	*c = 0
}

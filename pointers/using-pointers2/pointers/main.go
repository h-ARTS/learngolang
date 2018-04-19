package main

import (
	"fmt"
)

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is now 0, because of passing a value on the same memory address!
}

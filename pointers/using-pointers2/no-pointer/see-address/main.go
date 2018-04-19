package main

import (
	"fmt"
)

func zero(z int) {
	fmt.Printf("%p \n", &z) // address in func zero
	fmt.Println(&z)         // address in func zero
	fmt.Println(z)          // z is also 5 but has a different memory address
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p \n", &x) // address in func zero
	fmt.Println(&x)         // address in func zero
	zero(x)
	fmt.Println(x) // x is till 5
}

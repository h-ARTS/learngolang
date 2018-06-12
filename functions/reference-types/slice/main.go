package main

import (
	"fmt"
)

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Hanan]
}

func changeMe(z []string) {
	z[0] = "Hanan"
	fmt.Println(z) // [Hanan]
}

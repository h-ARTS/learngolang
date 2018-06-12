package main

import "fmt"

func main() {

	half := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}
	fmt.Println(half(2))
}

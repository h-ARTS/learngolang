package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", callback)
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs)
}

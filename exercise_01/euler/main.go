package main

import (
	"fmt"
	"math"
)

func sumSqr(c int) int {
	var n float64 = 1
	r := 0
	for i := c; i >= 1; i-- {
		r += int(math.Pow(n, 2))
		n++
	}
	return r
}

func sqrSum(c int) int {
	var r int
	for i := 0; c >= i; i++ {
		r += i
		if i == c {
			r = int(math.Pow(float64(r), 2))
		}
	}
	return r
}

func main() {
	fmt.Println(sqrSum(10) - sumSqr(10))
	fmt.Println(sqrSum(100) - sumSqr(100))
}

/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

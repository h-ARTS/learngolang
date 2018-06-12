package main

import (
	"fmt"
)

func greatest(n ...int) int {
	var largest int
	for _, v := range n {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	fmt.Println(greatest(2, 4, 1, 333, 22, 75, 234, 978, 12, 0, 300, 1200))
}

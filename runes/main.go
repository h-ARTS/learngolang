package main

import "fmt"

func main() {
	for i := 100; i <= 1800; i++ {
		fmt.Println(i, " – ", string(i), " – ", []byte(string(i)))
	}
}

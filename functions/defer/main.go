package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world!")
}

func main() {
	defer world() // remove defer and see what it does!
	hello()
	// defer gets executed right before the main() exits...
}

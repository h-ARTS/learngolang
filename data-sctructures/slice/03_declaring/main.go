package main

import (
	"fmt"
)

func main() {
	student := []string{}         // shorthand declartion
	var student2 []string         // var slice declaration
	teacher := make([]string, 35) // slice with make -> declaring the length and capacity
	teacher = append(teacher, "John Doe")

	fmt.Println(student)
	fmt.Println(student == nil)
	fmt.Println(student2)
	fmt.Println(student2 == nil)
	fmt.Println(teacher)
	fmt.Println(teacher == nil)
}

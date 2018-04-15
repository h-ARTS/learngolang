package main

import "fmt"

func main() {

	/* These are shorthand variables */
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`

	/* Note those formats:
	%v = default formatting for value
	%T = a Go-syntax representation of the type of the value */
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

}

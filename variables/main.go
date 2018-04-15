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

	// Zero-value variables, which will be assigned later in case...
	var ab int
	var ba string
	var cb float64
	var dc bool

	fmt.Printf("%v \n", ab)
	fmt.Printf("%v \n", ba)
	fmt.Printf("%v \n", cb)
	fmt.Printf("%v \n", dc)

	fmt.Println()

}

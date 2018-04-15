package main

import "fmt"

var text = "this is stored in the variable a"          // package scope
var btext, ctext string = "stored in b", "stored in c" // package scope
var dtext string                                       // package scope

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
	var ab int     // function scope
	var ba string  // function scope
	var cb float64 // function scope
	var dc bool    // function scope

	fmt.Printf("%v \n", ab)
	fmt.Printf("%v \n", ba)
	fmt.Printf("%v \n", cb)
	fmt.Printf("%v \n", dc)

	fmt.Println()

	// Declare as many as you want and then assign it later!
	var chatMessage string
	var k, l, m int
	k = 1

	chatMessage = "Hello World!"

	fmt.Println(chatMessage, k, l, m)

	// Init as many as you want!
	var message = "Hello World!"
	var x, y, z int = 1, 2, 3

	fmt.Println(message, x, y, z)

}

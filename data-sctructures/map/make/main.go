package main

import (
	"fmt"
)

func main() {

	// 2 ways of making a map:
	var myGreeting = make(map[string]string)
	var myCars = map[string]string{}

	myGreeting["Tim"] = "Good Morning"
	myGreeting["Mbabbe"] = "Bonjour"

	myCars["Audi"] = "A3 1.8L TFSI"
	myCars["Toyota"] = "Supra"

	fmt.Println(myGreeting)
	fmt.Println(myCars)
}

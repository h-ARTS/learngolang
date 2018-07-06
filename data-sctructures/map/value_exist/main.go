package main

import (
	"fmt"
)

func main() {

	myGreetings := map[int]string{
		0: "Guten Morgen!",
		1: "Good Morning!",
		2: "Assalaamu 'Aleikum!",
		3: "Buenos Dias!",
	}

	fmt.Println(myGreetings)
	// delete(myGreetings, 3)

	if val, exists := myGreetings[3]; exists {
		fmt.Println("That value exists.")
		fmt.Println("Value: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("Value: ", val)
		fmt.Println("Exists: ", exists)
	}

	fmt.Println(myGreetings)
}

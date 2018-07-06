package main

import (
	"fmt"
)

func main() {

	var myTrucks = map[string]string{
		"MAN":           "6 wheels",
		"Mercedes-Benz": "Sprinter",
	}
	// Add
	myTrucks["Renault"] = "Master"

	// Update
	myTrucks["Renault"] = "Scenic"

	// Delete
	delete(myTrucks, "Renault")

	fmt.Println(myTrucks)
}

package main

import "fmt"

func main() {
	for i := 100; i <= 180; i++ {
		fmt.Println(i, " – ", string(i), " – ", []byte(string(i)))
	}
	foo := 'a' // 'a' <-- rune "a" <-- string
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}

package main

import "fmt"

/* https://tour.golang.org/basics/11 */
func main() {
	//var name = "Edmilson"
	var age int32 = 23
	var isCool = true
	isCool = false

	// Shorthand
	size := 1.3

	name, email := "Edmilson", "ed@gmail.com"

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", size)
}

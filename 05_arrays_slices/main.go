package main

import "fmt"

func main() {
	//var fruitArray [2]string
	//fruitArray[0] = "Apple"
	//fruitArray[1] = "Orange"

	fruitArray := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}

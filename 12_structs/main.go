package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

func (self Person) greet() string {
	return "Hello, my name is " + self.firstName + " " + self.lastName + " and I am " + strconv.Itoa(self.age)
}

func (self *Person) hasBirthday() {
	self.age++
}

func (self *Person) getMarried(spouseLastName string) {
	if self.gender == "m" {
		return
	} else {
		self.lastName = spouseLastName
	}
}

func main() {
	/*person1 := Person{
	firstName: "Edmilson",
	lastName: "Santana",
	city:     "Recife",
	gender:   "M",
	age:      23} */
	person1 := Person{"Samantha", "Smith", "Recife", "f", 23}

	fmt.Println(person1.firstName)

	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())
}

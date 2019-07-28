package main

import "fmt"

func main() {
	//emails := make(map[string]string)
	emails := map[string]string{"ed": "ed@gmail.com", "Sharon": "sharon@gmail.com"}
	//emails["ed"] = "ed@gmail.com"
	//emails["cla"] = "cla@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["ed"])
	fmt.Println(len(emails))

	delete(emails, "mike")
	fmt.Println(emails)
}

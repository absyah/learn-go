package main

import "fmt"

func main() {
	// with data type
	var firstName string = "John"
	var lastName string
	lastName = "Doe"

	// without data type
	message := "Hello"

	// constant
	const planet string = "Earth"

	fmt.Printf("%s, %s %s \n", message, firstName, lastName)
	fmt.Printf("You're on %s", planet)
}
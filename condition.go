package main

import "fmt"

func main() {
	// if else
	var isPaid bool = false

	if (isPaid) {
		fmt.Println("Paid")
		return
	}

	fmt.Println("Unpaid")

	// switch case
	var selectedNumber int8 = 3

	switch selectedNumber {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Default")
	}
}
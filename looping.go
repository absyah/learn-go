package main

import "fmt"

func main() {
	// for
	for i := 0; i < 5; i++ {
		fmt.Println("Loop #1:", i)
	}

	var i int = 0

	for i < 5 {
		fmt.Println("Loop #2:", i)
		i++
	}


}
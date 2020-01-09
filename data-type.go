package main

import "fmt"

func main() {
	var string string = "Hello"
	var integer uint8 = 0
	var float float32 = 0.0
	var boolean bool  = true

	fmt.Printf("This is string '%s'\n", string)
	fmt.Printf("This is integer '%d'\n", integer)
	fmt.Printf("This is float '%3f'\n", float)
	fmt.Printf("This is boolean '%t'\n", boolean)
}
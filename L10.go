//write a program that converts strings to bool.
package main

import (
	"fmt"
)

func L10() {
	input := "1"
	var output bool

	switch input {
	case "1", "yes", "true":
		output = true
	case "2", "no", "0":
		output = false
	default:
		fmt.Println("That is not a valid response.")
	}
	fmt.Println(output)
}

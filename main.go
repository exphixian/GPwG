package main

import (
	"fmt"
)

func main() {
	lesson := ""
	fmt.Println("Which lesson are you referring to?")
	fmt.Println("Lessons 5, 6, & 7 are currently supported.")
	fmt.Scan(&lesson)

	//test input
	//fmt.Println(lesson)

	switch lesson {
	case "5":
		L5()
	case "6":
		L6()
	case "6.5":
		L6Alt()
	case "7":
		L7()
	default:
		fmt.Println("That is not a valid lesson. Ending program.")
	}

}

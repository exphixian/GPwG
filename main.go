package main

import (
	"fmt"
)

func main() {
	lesson := ""
	fmt.Println("Which lesson are you referring to?")
	fmt.Println("Lessons 2 - 9 are currently included.")
	fmt.Scan(&lesson)

	//test input
	//fmt.Println(lesson)

	switch lesson {
	case "2":
		L2()
	case "3":
		L3()
	case "5":
		L5()
	case "6":
		L6()
	case "6.5":
		L6Alt()
	case "7":
		L7()
	case "8":
		L8()
	case "9":
		L9()
	default:
		fmt.Println("That is not a valid lesson. Ending program.")
	}

}

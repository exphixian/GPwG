package main

import (
	"fmt"
)

func main() {
	lesson := ""
	fmt.Println("Which lesson are you referring to?")
	fmt.Println("Lessons 2 - 15 are currently complete.")
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
	case "10":
		L10()
	case "11":
		part := ""
		fmt.Println("Part 1 or 2?")
		fmt.Scan(&part)
		if part == "1" {
			L11Decrypt()
		} else if part == "2" {
			L11Encrypt()
		} else {
			fmt.Println("That is not a valid lesson. Ending program.")
		}
	case "12":
		L12()
	case "13":
		L13()
	case "14":
		L14()
	case "15":
		L15()
	default:
		fmt.Println("That is not a valid lesson. Ending program.")
	}

}

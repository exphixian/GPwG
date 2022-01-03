//write a guess the number program. make the computer pick random numbers between 1 & 100 until it guesses your number, which you declare at the top of the program.
//Display each guess and whether it was too big or too small.

package main

import (
	"fmt"
	"math/rand"
)

func L3() {
	var number, guess int
	correct := false
	fmt.Println("Pick a number between 1 & 100:")
	fmt.Scan(&number)

	for !correct {
		guess = rand.Intn(100) + 1
		//fmt.Println(guess)
		switch {
		case guess == number:
			correct = true
			fmt.Printf("\nThe number is %v!\n", guess)
		case guess < number:
			fmt.Printf("%v is too small to be your number... \n", guess)
		case guess > number:
			fmt.Printf("Whoops, %v is too big!\n", guess)
		}

	}
}

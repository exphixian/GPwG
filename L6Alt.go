package main

import (
	"fmt"
	"math/rand"
)

const nickel, dime, quarter = .05, .1, .25

func L6Alt() {
	piggybank := 0.0
	var totalnickels, totaldimes, totalquarters int = 0, 0, 0
	for piggybank <= 20 {
		switch i := rand.Intn(3); i {
		case 0:
			piggybank += nickel
			totalnickels++
		case 1:
			piggybank += dime
			totaldimes++
		case 2:
			piggybank += quarter
			totalquarters++
		}
		fmt.Printf("You have %4.2f in your piggy bank - %v quarters, %v dimes, and %v nickels.\n", piggybank, totalquarters, totaldimes, totalnickels)

	}
}

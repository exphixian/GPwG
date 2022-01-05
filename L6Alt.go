package main

import (
	"fmt"
	"math/rand"
)

func L6Alt() {
	piggybank := 0.0
	nickels, dimes, quarters := 0, 0, 0
	for piggybank <= 20 {
		switch i := rand.Intn(3); i {
		case 0:
			piggybank += .05
			nickels++
		case 1:
			piggybank += .1
			dimes++
		case 2:
			piggybank += .25
			quarters++
		}
		fmt.Printf("You have %4.2f in your piggy bank - %v quarters, %v dimes, and %v nickels.\n", piggybank, quarters, dimes, nickels)

	}
}

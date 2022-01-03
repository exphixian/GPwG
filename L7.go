package main

import (
	"fmt"
	"math/rand"
)

func L7() {
	piggybank := 0

	for piggybank <= 2000 {
		switch i := rand.Intn(3); i {
		case 0:
			piggybank += 5
		case 1:
			piggybank += 10
		case 2:
			piggybank += 25
		}
		dollars := piggybank / 100
		cents := piggybank % 100

		fmt.Printf("$%d.%02d\n", dollars, cents)
	}

}

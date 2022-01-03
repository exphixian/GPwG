package main

import (
	"fmt"
	"math/rand"
)

func L6() {
	value := 0.0

	for value <= 20.0 {
		coin := 0.05
		switch i := rand.Intn(3); i {
		case 0:
			coin = 0.10
		case 1:
			coin = 0.25
		}
		value += coin

		fmt.Printf("%4.2f\n", value)
	}
}

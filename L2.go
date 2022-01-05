// Get Programming with Go page 22 prompt:
// Write a program to determine how fast a ship would need to travel in km/h in order to reach Mars in 28 days. Assume a distance of 56,000,000 km.

package main

import "fmt"

func L2() {
	const hoursPerDay = 24
	var distance = 56000000

	speed := distance / (28 * hoursPerDay)

	fmt.Printf("You would need to travel %v km/h to reach Mars in 28 days.\n", speed)

}

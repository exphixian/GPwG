package main

import (
	"fmt"
	"math/rand"
)

//const deptarture = "10/13/20"
const distance = 62100000

func L5() {
	fmt.Println("Spaceline			Duration	Roundtrip	Price")
	fmt.Println("=============================================================================")

	//generate 10 tickets
	for ticket := 0; ticket < 10; ticket++ {

		//October 13th 2020 is departure date - 62,100,000 km is distance. Speed will be 16-30 km/s.
		speed := rand.Intn(15) + 17
		traveltime := distance / (speed * 86400)
		price := rand.Intn(5) + 36

		//Faster ships are more expensive - price should range between 36 & 56 m. x2 for round trip
		if speed > 25 {
			price = price + 15
		} else if speed > 20 {
			price = price + 10
		}

		roundTrip := ("One-way   ")
		switch i := rand.Intn(2); i {
		case 0:
			roundTrip = ("Round Trip")
			price = price * 2

		}

		//randomly select the spaceline (SpaceX, Space Adventures, or Virgin Galactic)
		spaceline := "SpaceX          "
		switch i := rand.Intn(3); i {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "Virgin Galactic "

		}

		/*display in a tabular format with 4 columns:
		* company providing the service
		* duration in days (1 way)
		* whether the price includes a round trip
		* price in millions */
		fmt.Printf("%s		%d days		%s	%d mil USD\n", spaceline, traveltime, roundTrip, price)
	}

}

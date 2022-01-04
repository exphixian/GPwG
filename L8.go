/*convert 236e15km to light years
lightspeed := 299792 km/sec
secondsPerDay := 86400*/

package main

import (
	"fmt"
)

const lightspeed, secondsPerDay, distance = 299792, 86400, 236000000000000000

func L8() {
	fmt.Println(distance / (lightspeed * secondsPerDay * 365))
}

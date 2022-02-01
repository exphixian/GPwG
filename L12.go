package main

import (
	"fmt"
)

//converts Kelvin to Celsius
func k2c(k float64) float64 {
	k -= 273.15
	return k
}

func cel2feh(c float64) float64 {
	c = (c * 9.0 / 5.0) + 32.0
	return c
}

func k2f(k float64) float64 {
	k = ((k - 273.15) * 9.0 / 5.0) + 32.0
	return k
}

func L12() {
	kelvin := 294.0
	celsius := k2c(kelvin)
	fmt.Printf("%4.2f C\n", celsius)
	fahrenheit := cel2feh(celsius)
	fmt.Printf("%4.2f F\n", fahrenheit)
	kelvin = 0
	kel2Far := k2f(kelvin)
	fmt.Printf("If K is 0, F is %4.2f\n", kel2Far)
}

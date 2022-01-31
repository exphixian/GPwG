// write a program with C, F, & K types and methods to convert between the three
package main

import (
	"fmt"
)

type celsius float64

//converts C to F
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

//converts C to K
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

type fahrenheit float64

//converts F to C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

//converts F to K
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

//uncomment below to use with main
//type kelvin float64

//converts K to C
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

//converts K to F
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func L13() {
	var k kelvin = 294.0
	c := k.celsius()
	f := k.celsius().fahrenheit()
	fmt.Printf("%5.2f K is %5.2f C & %5.2f F.\n", k, c, f)
}

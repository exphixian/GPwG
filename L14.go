package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	rand.Seed(time.Now().UnixNano())
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func L14() {
	var offset kelvin = 5
	samples := 6

	for i := 0; samples > i; i++ {
		sensor := calibrate(fakeSensor, offset)
		fmt.Println(sensor())
		time.Sleep(time.Second)
	}

}

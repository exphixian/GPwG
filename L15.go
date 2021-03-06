package main

import (
	"fmt"
	"time"
)

func c2f(c int) int {
	return (c * 9 / 5) + 32
}

func f2c(f int) int {
	return (f - 32) * 5 / 9
}

func drawTable(tempType1 string, tempType2 string, function func(int) int) {
	row := "================================="
	fmt.Println(row)
	fmt.Printf("|	%v	|	%v	|\n", tempType1, tempType2)
	fmt.Println(row)
	for i := -40; i < 100; i += 5 {
		adjTemp := function(i)
		fmt.Printf("|	%v	|	%v	|\n", i, adjTemp)
	}
	fmt.Println(row)

}

func L15() {
	drawTable("C", "F", c2f)
	time.Sleep(time.Second)
	fmt.Println("")
	time.Sleep(time.Second)
	drawTable("F", "C", f2c)
}

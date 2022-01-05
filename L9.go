//deciper a quote.
package main

import (
	"fmt"
)

func L9() {
	quote := "L fdph, L vdz, L frqtxhuhg."
	fmt.Printf("Part 1 - decipher a Caesar cipher: %v\n", quote)

	for i := 0; i < len(quote); i++ {
		c := quote[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("\n")

	message := "Hola Estacion Espacial Internacional"
	fmt.Printf("Part 2 - Ciper message with ROT13: %v\n", message)

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println("")
}

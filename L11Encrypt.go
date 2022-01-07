package main

import (
	"fmt"
	"strings"
)

func L11Encrypt() {
	// c > C = 32
	plainText := "Your message goes here"
	key := "GOLANG"

	plainText = strings.ToUpper(strings.Replace(plainText, " ", "", -1))
	key = strings.ToUpper(strings.Replace(key, " ", "", -1))

	for i := 0; i < len(plainText); i++ {
		c := plainText[i] - 'A'
		keyIndex := i % len(key)
		k := key[keyIndex] - 'A'
		cipher := c + k + 'A'
		if cipher > 90 {
			cipher -= 26
		} else if cipher < 64 {
			cipher += 26
		}
		fmt.Printf("%c", cipher)
	}
	fmt.Println()
}

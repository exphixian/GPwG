//write a program to decipher text
package main

import (
	"fmt"
)

func L11Decrypt() {
	cipher := "CSOITEUIWUIZNSROCNKFD"
	key := "GOLANG"

	for i := 0; i < len(cipher); i++ {
		c := cipher[i] - 'A'
		keyIndex := i % len(key)
		k := key[keyIndex] - 'A'
		output := c - k + 'A'
		if output < 64 {
			output += 26
		} else if output > 90 {
			output -= 26
		}
		fmt.Printf("%c", output)
	}
	fmt.Println("")

}

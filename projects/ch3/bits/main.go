package main

import (
	"fmt"
	"main/bitlib"
)

func main() {
	var result uint8
	fmt.Println("bits demo")

	// clear all 4 bits using 15
	result = bitlib.BitClear(15, 15)
	// clear last two bits using 3
	result = bitlib.BitClear(15, 3)
	fmt.Printf("result: %x\n", result)
}

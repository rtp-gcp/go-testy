package popcount

// package main

import (
	"fmt"
)

// pc[i] is the population count of i
// population count is number of 1's in a byte.
// byte is the type.  Its an 8bit unsigned integer.
var pc [256]byte

// This is a special init() function which is called automatically
func init() {
	// iterate ofver all indices of the array (0-255)
	// By default, it returns the index and the value for
	// an array.  Ignoring the value could be done like so:
	// for i, _ := range pc {...}
	for i := range pc {
		// That is a neat trick to count the number of ones in a byte.
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits ) of 64-bit value x
func PopCount(x uint64) int {
	// one way to define, but generates a warning
	// var result int = 0
	result := 0
	var byteValue int

	// Debug input
	fmt.Printf("x: %d  %x\n", x, x)

	for i := 0; i < 8; i++ {

		// Could mask so we have one byte
		// byteValue = int(0xFF & (x >> (i * 8)))

		// Or use types to mask for us
		byteValue = int(byte(x >> (i * 8)))

		fmt.Printf("i:%d:      x>>i: %x\n", i, byteValue)
		result = result + int(pc[byteValue])
	}

	return result
}

func main() {
	var result int
	result = PopCount(15)
	fmt.Println(result)
}

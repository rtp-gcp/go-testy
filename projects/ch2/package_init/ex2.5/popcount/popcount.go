package popcount

// package main
//
// Exercis e 2.5: The expression x&(x-1) clears the rig htmost non-zero bit of x.
// Write a version of PopCountthat counts bits by using this fac t, and ass ess its per for mance.

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

	// Debug input
	fmt.Printf("x(decimal,hex,binary): %d  0x%04x \t %064b\n", x, x, x)

	for x != 0 {

		// mask so we have one bit
		x = x & (x - 1)

		fmt.Printf("x(decimal,hex,binary): %d 0x%04x \t %064b\n", x, x, x)
		result = result + 1
	}

	return result
}

func main() {
	var result int
	result = PopCount(15)
	fmt.Println(result)
}

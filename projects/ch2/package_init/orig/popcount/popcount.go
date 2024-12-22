package popcount

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
	// This code splits each byte of the 64-bit value into bytes.
	// For each byte, it returns the precomputed value from the lookup
	// table to get the number of bits for that byte.  It then sums up the
	// counts.
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

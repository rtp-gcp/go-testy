// various temp utils
package main

import (
	"fmt"
)

const boilingF = 212.0

// so unicode
// F degree symbol is \u2109
// C degree symbol is \u2103

func boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g\u2109 or %g\u2103\n", f, c)
}

func boiling2(f float64) float64 {
	return (f - 32) * 5 / 9
}

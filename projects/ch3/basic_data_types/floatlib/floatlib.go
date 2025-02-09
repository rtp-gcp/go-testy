package floatlib

import (
	"fmt"
	"math"
)

func SomeFunc() int {
	return 0
}

func DemoFloats() int {
	const e = 2.71828 //

	// chatgpt constant
	// 6.02214076 × 10²³
	// const Avogadro = 6.02214076e23
	// book constant
	const Avogadro = 6.02214129e23
	const Planck = 6.62606957e-34

	// Using old c style print format specifier
	fmt.Printf("----------------------\n")
	fmt.Printf("e: %8.3f\n", e)
	fmt.Printf("Avogadro: %f\n", Avogadro)
	fmt.Printf("Avogadro: %8.3f\n", Avogadro)
	fmt.Printf("Planck: %8.3f\n", Planck)
	fmt.Printf("\n")

	// Using golang g  specifier
	fmt.Printf("----------------------\n")
	fmt.Printf("e: %8.3g\n", e)
	fmt.Printf("Avogadro: %8.3g\n", Avogadro)
	fmt.Printf("Planck: %8.3g\n", Planck)

	for x := 0; x < 8; x++ {
		// super script 2, type ctrl + V u00B3
		// super script x, type ctrl + V u02E3
		fmt.Printf("x = %d     eˣ = %8.3f\n", x, math.Exp(float64(x)))
	}

	return 0
}

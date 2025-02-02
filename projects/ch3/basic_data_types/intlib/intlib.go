package intlib

import "fmt"

func SomeFunc() int {
	return 0
}

func DemoRevForLoop() {
	// len returns a signed int so that it can be used
	// to make reverse loops by subtraction
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}

func DemoForTypeConversion() {
	var apples int32 = 1
	var oranges int16 = 2

	// This is an error because they are different types
	// var compote int = apples + oranges

	// TThis works
	// var compote int = int(apples) + int(oranges)
	// This also works and removes the int type for the variable since we have typecast
	// compote := int(apples) + int(oranges)
	compote := int(int16(apples) + oranges)

	var myPi float64 = 3.14
	var myBigFloat float64 = 123456.789
	var o int = 333
	var w int = 4095
	fmt.Printf("float type cast of pi to int: %d\n", int(myPi))
	// This demos using two positional parms twice. #[1] where 1 is first postitional parameter
	fmt.Printf("Demo using multi params %d %d\n", 1, 2)
	// The o in the print specifier is octal. its not the variable.
	fmt.Printf("Demo using multi params %[1]o %[1]o\n", o)
	// here we show its hex using x
	fmt.Printf("Demo using multi params %[1]x %[1]x\n", w)
	// now do with two parms
	fmt.Printf("Demo using multi params %[2]d %[1]d\n", 1, 2)
	// now try with one that is typedef
	// The integer part of 123456.789 in hexadecimal is 0x1e240.
	fmt.Printf("Demo using typecast arg %[1]f %[1]d %[1]x\n", int(myBigFloat))

	fmt.Println(compote)
}

func DemoForRunes() {
	// demo for runes golang here
	//
	// A rune in Go is an alias for int32 and is used to represent a Unicode code point.
	// Since Go uses UTF-8 encoding, runes are useful when dealing with
	// characters beyond standard ASCII (e.g., emojis, special symbols, non-English letters).

	fmt.Println("Demo for rune")

	var r rune = 'A' // Single quotes are used for runes

	// Unicode representation
	fmt.Printf("Rune: %c\n", r)               // Output: A
	fmt.Printf("Unicode Code Point: %U\n", r) // Output: U+0041

	// Iterate over a string (UTF-8 encoded)
	str := "Go语言" // Contains Chinese characters
	for i, r := range str {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
	}

	// Rune slice (array of runes)
	runes := []rune("Hello, 世界")              // Mixed English and Chinese
	fmt.Println("Rune Slice:", string(runes)) // Convert back to string
}

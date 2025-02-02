package main

import (
	"fmt"
	"main/bitlib"
	"main/floatlib"
	"main/intlib"
)

func main() {
	var result uint8
	fmt.Println("bits demo")

	// clear all 4 bits using 15
	result = bitlib.BitClear(15, 15)
	// clear last two bits using 3
	result = bitlib.BitClear(15, 3)
	fmt.Printf("result: %x\n", result)

	intlib.DemoRevForLoop()
	floatlib.SomeFunc()
	intlib.SomeFunc()
	intlib.DemoForTypeConversion()
}

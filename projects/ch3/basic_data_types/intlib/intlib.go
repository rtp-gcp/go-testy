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
	compote := int(apples) + int(oranges)

	fmt.Println(compote)
}

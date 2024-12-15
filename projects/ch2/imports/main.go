package main

import (
	"fmt"
	"main/tempconv"
	"os"
	"strconv"
)

func main() {
	// fmt.Printf("arg[0]: %v\n", os.Args[0])
	if len(os.Args) == 1 {

		fmt.Printf("USAGE: ./main <nn.nn> \n")
		fmt.Printf("USAGE: specify a number for a temperature.\n")
		return
	}

	for _, arg := range os.Args[1:] {

		// fmt.Printf("arg: %v\n", arg)

		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("As Fahrenheit: %s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("As Celsius: %s = %s\n", c, tempconv.CToF(c))
	}
}

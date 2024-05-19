// prints command line arguments
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// These all define a string
	// s:= ""
	// var s string
	// var s = ""
	// var s string = ""
	var s, sep string

	// the parens are never used inn the for loop
	// init; condition; post line.
	// If the triplet line is not specified, its
	// an infinite loop.
	for i := 1; i < len(os.Args); i++ {
		// In this case + is string concatenation
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("args: ", s)

	// alternative iterating over a range
	// in this case we don't need the index, just the arg so
	// use the "blank identifier".  Use this when compiler needs
	// a variable but the program logic does not.
	//
	s = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	fmt.Println("args: ", s)

	// Each time above, the string gets
	// use the "blank identifier".  Use this when compiler needs
	// a variable but the program logic does not.
	//
	// Just use the strings.Join() function to do the same thing
	// without looping or creating multiple strings.
	fmt.Println("args: ", strings.Join(os.Args[1:], "  "))

	// Ex1.1
	fmt.Println("program name: ", os.Args[0])

	// Ex1.2
	// assumes 10 args at most
	s = ""
	var args [10]string
	for index, arg := range os.Args[1:] {
		args[index] = arg
	}
	for i := 0; i <= len(os.Args[1:]); i++ {
		fmt.Println("index: ", i, "    arg: ", args[i])

	}

	// Ex1.3
	// measure running time
	start := time.Now()
	s = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
	}
	secs := time.Since(start).Seconds()
	fmt.Println("running time: ", secs)
}

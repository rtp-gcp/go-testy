// finds dupelines
// ch 1.3
// dupelines prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// built-in function make() creates an empty map
	counts := make(map[string]int)
	// The bufio package helps with input and output io.
	// Scanner reads lines or words.
	// It removes the newline from the line.
	// It returns
	//    true if there is a line of input read
	//    false if no input to read
	input := bufio.NewScanner(os.Stdin)

	// each time it reads a line, it uses the line
	// as a key in the map, and the value is the number
	// of times the line is encountered.
	// The key is the line, the value is incremented.
	for input.Scan() {
		counts[input.Text()]++
		// above is equivalent to
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}
	// first time a line is found and a key is created,
	// the value is set to zero.

	// NOTE: ignoring potential errors from input.Err()
	//
	// This is a range based for loop, iterating over the map.
	// Each iteration gives a key and a value.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

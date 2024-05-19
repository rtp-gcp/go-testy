// Lissajous generates GIF animations of random Lissajous figures
// ch 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

type fileDetails struct {
	fileName  string
	dupeLines int
}

func main() {

	// built-in function make() creates an empty map
	// make returns a refernce.
	counts := make(map[string]int)
	var fileName string
	fileInfo := make(map[string]fileDetails)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "none", nil)
	} else {
		for _, arg := range files {
			// op.Open returns two values:
			// file handle and status.
			// When status is nil, open is ok.
			// note the status is of type error,
			// hence when no error occurs its nil.
			// when an error does occur the value
			// corresponds to the type of error.
			// think of an enum.
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// In golang, you can put function calls and function
			// declarations in any order.  So here we are calling
			// the function before we have defined it.
			fileName = arg
			countLines(f, counts, fileName, fileInfo)
			f.Close()

		}
	}
	// This is a range based for loop, iterating over the map.
	// Each iteration gives a key and a value.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", fileName, n, line)
		}
	}

	// EX 1.4
	for key, value := range fileInfo {
		fmt.Println(key, value)
	}
}

// counts is a variable of type map. A map is a reference
// to the data structure created by make.
// When a map is passed to a function, the function
// receives a copy of the reference, so any changes
// the called function makes to the underlying
// data structure will be visible through the caller's
// map reference too.  This means any changes to the
// counts in the map in the function are seen by main.
//
// Ex 1.4 we pass in the file name and the info map
func countLines(f *os.File, counts map[string]int, fileName string, fileInfo map[string]fileDetails) {

	// The bufio package helps with input and output io.
	// Scanner reads lines or words.
	// It removes the newline from the line.
	// It returns
	//    true if there is a line of input read
	//    false if no input to read
	input := bufio.NewScanner(f)
	for input.Scan() {
		// each time it reads a line, it uses the line
		// as a key in the map, and the value is the number
		// of times the line is encountered.
		// The key is the line, the value is incremented.
		counts[input.Text()]++
		fileInfo[fileName] = fileDetails{fileName: fileName, dupeLines: counts[input.Text()]}
	}
	// NOTE: ignorning potential errors from input.Err()
	// Ex 1.4
}

// Ex1.4
// Modify dup2 to print the names of all
// files in which each duplicated line occurs

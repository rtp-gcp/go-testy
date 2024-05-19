// finds dupelines in files rather than stdin
// ch 1.3
// dupelines prints the text of each line that appears more than
// once in a file, preceded by its count.
//
// Dupelines2 worked by streaming the files and processing them
// line by line, file by file.
// In this version, an entire file is read and then the process
// of determining the duplicate lines is performed.
//
// This variant introduces:
//   - ReadFile() from io/ioutil package which reads the entire file
//     into memory.
//   - strings.Split which splits a string into a slice of substrings
//   - strings.Split is opposite of strings.Join
package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	// built-in function make() creates an empty map
	// make returns a refernce.
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile returns a byte slice that must be converted
		// into a string so it can be split by strings.Split.
		//
		// Under the covers,
		//  * bufio.Scanner
		//  * ioutil.ReadFile
		//  * ioutil.WriteFile
		// all use the Read/Write methods of *os.File
		// but, its rare to use these lower level routines
		// directly.  The higher level routines: Scanner,
		// Read/WriteFile are easier to use.
		data, err := ioutil.ReadFile(filename)
		// In vscode, you can type 'iferr' and then
		// select the autocomplete and it will auto
		// type this type of stanza.
		// Likewise, for a for range loop, type 'forr'
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

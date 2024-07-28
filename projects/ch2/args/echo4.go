// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// This packages provides help out of the box

// DEMO
// $ ./echo4 foo
// foo
// $ ./echo4 -s ":" foo goo
// foo:goo
// $ ./echo4 --help
// Usage of ./echo4:
//  -n    omit trailing newline
//  -s string
//        separator (default " ")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}

}

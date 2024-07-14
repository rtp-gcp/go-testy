package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println("foo: ", foo)
	boiling()
	fmt.Printf("100F is %g\n", boiling2(100))
}

var foo int = 22

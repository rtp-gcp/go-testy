# go testing

Typically test code can not use code in package main.  Add component code to a subdir/package.  
Afterwards add packagename_test.go in that directory for the test code.

```
$ cd ex2.3
$ og mod tidy  --- to pull in the testing framework and the popcount package
$ go test -v ./...    - to recursively run all tests in subdirectories
```



## Dir setup

```
main.go 
bitlib/bitlib.go
bitlib/bitlib_test.go 
```

## Contents of main.go 

```
package main

import (
	"fmt"
	"main/bitlib"
)


func main() {
	var result uint8
	fmt.Println("bits demo")
	result = bitlib.BitClear(15)
	fmt.Printf("result: %x\n", result)
}
```

## contents of bitlib/bitlib.go 

```
package bitlib

func BitClear(value uint8) uint8 {
	return value
}
```

## contents of bitlib/bitlib_test.go 

```
package bitlib

import (
	"testing"
)

func TestAnd(t *testing.T) {
	result := BitClear(15)
	if result != 15 {
		t.Error("BIT AND fail")
	}
}

func TestClear(t *testing.T) {
	result := BitClear(15)
	if result != 15 {
		t.Error("BIT CLEAR fail")
	}
}
```

## to run all tests 

From the main directory
```
$ cd xxx/main
$ go test -v ./...
```

From the package directory
```
$ cd xxx/main/bitlib
$ go test -v 
```

## to run a specific test

From the package directory
```
$ cd xxx/main/bitlib
```

For just a specific test name

```
go test -run BitClear
```

For a regex expression to find test

```
$ go test -v -run ^TestClear$
# escape zero or more asterisk by shell
$ go test -v -run ^.\*Clear$
```

## to benchmark a test

Need to add some test code to call a function multiple times:

```
func BenchmarkFunction(b *testing.B) {
        for i := 0; i < b.N; i++ {
                popcount.PopCount(64 * 1024)
        }
}
```

This will run a specific test and then call the benchmark test
```
$ cd popcount
$ go test -v -run ^TestFunction64k$ -bench=BenchmarkFunction
```

package main

import (
	"fmt"
	"main/popcount"
)

func main() {
	fmt.Println(popcount.PopCount(15))   // Output 4 (binary: 0x1111)
	fmt.Println(popcount.PopCount(255))  // Output 8 (binary: 0x1111 1111)
	fmt.Println(popcount.PopCount(256))  // Output 1 (binary: 0x1000 0000)
	fmt.Println(popcount.PopCount(1024)) // Output 1 (binary: 0x1000 0000)
	// fmt.Println(popcount.PopCount(4095))        // Output 1 (binary: 0x1000 0000)
	// fmt.Println(popcount.PopCount(4096))        // Output 1 (binary: 0x1000 0000)
	// fmt.Println(popcount.PopCount(64 * 1024))   // Output 1 (binary: 0x1000 0000)
	fmt.Println(popcount.PopCount(64*1024 - 1)) // Output 1 (binary: 0x1000 0000)
}

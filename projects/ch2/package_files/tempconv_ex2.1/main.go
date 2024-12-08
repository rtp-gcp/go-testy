package main

import (
  "fmt"
  //sincde this is in a module in a subdirectory, specify this package and the subdir
  tempconv "main/tempconv" 
)

// shift enter will select the auto suggestion
//
// To do this correctly
// 1. In this dir
//  - go mod init main
// 2. go mod tidy
// 3. go build
// 4. ./main
func main() {

  fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"
  fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
 
  // kelvin exercise
  fmt.Printf("zero celsius is  %v kelvin\n", tempconv.CToK(0) )
  fmt.Printf("zero kelvin is %v celsius\n",tempconv.KToC(0))
}

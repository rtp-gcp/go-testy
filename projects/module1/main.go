// A module test
//
// Demos:
//  * multiple files in one dir
//
//
// Steps
// In terminal do:
// 1. $ go mod init module1
//   -- This style is global namespace
//   -- demo did:
//      $ go mod init news.com/events
//      -- This style is using a non global namespace
//         The news.com/events is a virtual path.  The host could be github.com/netskink
//         and the app foo corresponding to the dir name of the git repo.
//   -- this creates:
//      go.mod which contains:
//         ```
//         module module1
//
//         go 1.20
//         ```
//       And says to run go mod tidy
// 2. go mod tidy
//   -- did nothing as far as I can tell.  go.mod is not modified
// 3. go run .
//   -- demos PublicName in second source file is visible
//   -- demos privateName in second source file is visible
//       -- this is unexpected
//       -- go vet does not say anything either
// 4. mkdir sub directory named subdir
// 5. Add new file (subtest.go) and put in package subdir

package main

// NOTE:
// When we did the module init, we specified this package main
// as part of module name "module1".  So, the subfolder package
// import starts with module1 and then the subdir name.
//
// NOTE:
// The import has two parts, the module name which is virtual and the
// subdirectory name.  In the example, he uses a go mod init new.com/events
// and so, he would have an equivalent line of:
// import news.com/events <virtual part> /subdir <real part>
//
// NOTE:
// To further obfuscate things, he demos that the package name does not need
// to match the directory name as shown here.  The import is subdir corresponding
// to the sub directory name, but the package in the subtest.go file is subdir1.
// Consquently, the usage in the main code is subdir1.PublicName.
//
// NOTE:
// the private name in lowercase letters in the subdir1 package is not
// visible in this package.
import (
	"fmt"
	subdir1 "module1/subdir"
)

func main() {
	fmt.Println("== main() ==")
	fmt.Println("main::PublicName: ", PublicName)
	fmt.Println("main::privateName: ", privateName)
	fmt.Println("subdir::PublicName: ", subdir1.PublicName)
	// This one is not visible
	// $ go vet .
	// will show this variable is not visible in this package.
	//fmt.Println("subdir::privateName: ", subdir.privateName)
}

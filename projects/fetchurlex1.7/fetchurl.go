// Fetch prints the content found at each specified URL.
// ch 1.4
// exercise 1.7 The function call to io.Copy(dst, src) reads from src and
// writes to dst.  Use it instead of iotuil.ReadAll to copy the response
// body to os.Stdout without requiring a buffer large enough
// to hold the entire stream.  Be sure to check the error result
// of io.Copy.
package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("needs an argument for url.")
		fmt.Printf("USAGE: %v <some url>\n", os.Args[0])
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		resp.Body.Close() // close the stream to avoid leak
	}
	os.Exit(0)
}

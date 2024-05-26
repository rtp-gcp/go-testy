// Fetch prints the content found at each specified URL.
// ch 1.5
// exercise 1.9
// Modify fetch to also print the HTTP status code, found in resp.Status.

package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("needs an argument for url.")
		fmt.Printf("USAGE: %v <some url>\n", os.Args[0])
		os.Exit(1)
	}

	// They can do multiple urls at a time
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			//fmt.Fprintf(os.Stderr, "missing prefix: %v\n", url)
			//os.Exit(1)
			url = "http://" + url
		}

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

		// Exercise 1.8 mod
		fmt.Fprintf(os.Stdout, "\n-------------------\n")
		fmt.Fprintf(os.Stdout, resp.Status)
		fmt.Fprintf(os.Stdout, "\n-------------------\n")
		fmt.Fprintf(os.Stdout, "%d\n", resp.StatusCode)
	}
	os.Exit(0)
}

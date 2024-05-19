// Fetch prints the content found at each specified URL.
// ch 1.5
package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close() // close the stream to avoid leak
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// write response to stdout
		fmt.Printf("%s", b)
	}
}

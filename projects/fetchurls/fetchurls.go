// Demos concurrency and channels
// ch 1.6

package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// By default?, main creates a goroutine.  In main, when
// the go fetch() statement is executed, it creates additional
// goroutines.
func main() {

	if len(os.Args) == 1 {
		fmt.Println("needs an argument for url.")
		fmt.Printf("USAGE: %v <some url>\n", os.Args[0])
		os.Exit(1)
	}

	start := time.Now()
	// Create a channel of string type each time!!!!
	ch := make(chan string)

	// They can do multiple urls at a time
	for _, url := range os.Args[1:] {
		// Create a new goroutine here !!!
		go fetch(url, ch) // start a goroutine with the newly created channel
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	os.Exit(0)
}

// Interesting, they are not using a return value.  Instead
// the code uses the channel parameter to return results.
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// ORIG ch <- fmt.Sprintf(err) // send error to channel ch
		ch <- fmt.Sprintf("error: %s\n", err.Error()) // send (formatted string method,) error to channel ch
		return
	}

	// copies the response body text to a null buffer using
	// ioutil.Discard.  Why do something so trivial?  Well,
	// its because we want to get the number of bytes of response
	// and this is a simple way to do so I suppose.
	// ioutil.Discard is used as the output stream.
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d %s", secs, nbytes, url)
}

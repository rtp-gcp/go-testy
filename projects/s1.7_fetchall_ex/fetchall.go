// Demos concurrency and channels
// ch 1.6
// EX 1.10 and 1.11 modify to:
// * save result
// * abort on timeout
//    - Use this url to simulate a long running (10s) fetch
//    - https://httpbin.org/delay/10

package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	// Create a context with a 3s timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Ensure the context is cancelled when main returns
	defer cancel()

	// Create a channel for messages
	ch := make(chan string)

	// They can do multiple urls at a time
	for index, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			//fmt.Fprintf(os.Stderr, "missing prefix: %v\n", url)
			//os.Exit(1)
			url = "http://" + url
		}
		// Start the goroutine
		go fetch(ctx, url, ch, index)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	os.Exit(0)
}

// Interesting, they are not using a return value.  Instead
// the code uses the channel parameter to return results.
func fetch(ctx context.Context, url string, ch chan<- string, counter int) {
	var filename string

	start := time.Now()

	// Create a new context with the previous context for the http.get so it can be cancelled.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("creating request for %s: %v", url, err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// ORIG ch <- fmt.Sprintf(err) // send error to channel ch
		ch <- fmt.Sprintf("error: %s\n", err.Error()) // send (formatted string method,) error to channel ch
		return
	}
	defer resp.Body.Close()

	//var body_as_text string
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//
	// write the contents to a file with a suffix of date and time
	//

	// Determine a filename based upon url
	// Just do a url counter like file1, file2, etc.
	filename = fmt.Sprintf("file%d-", counter)
	writeTextToFile(bytes, filename)

	// for some reason, when we did the copy, the body content was
	// was discarded in the source as well.
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d %s", secs, nbytes, url)

	resp.Body.Close() // don't leak resources

}

func writeTextToFile(bytes []byte, filename string) {

	// Get the current time in UTC
	currentTime := time.Now().UTC()

	// Format the time as a string
	//formattedTime := currentTime.Format("2006-01-02Z15:04:05 MST")
	// Format the time as ISO 8601 format
	formattedTime := currentTime.Format(time.RFC3339)

	// Print the formatted time
	//fmt.Println("Current time and date in New York:", formattedTime)

	// append the timestamp to the filename
	filename = filename + formattedTime

	// write content to file with specified filename
	err := os.WriteFile(filename, bytes, 0644)
	if err != nil {
		fmt.Println("error writing file")
	}
}

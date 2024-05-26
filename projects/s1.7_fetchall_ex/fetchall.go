// Demos concurrency and channels
// ch 1.6
// EX 1.10 and 1.11 modify to:
// * save result
// * abort on timeout

package main

// When using the intellisense for a routine, it will
// autodetect the package import and add it here.
import (
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

	fmt.Println("yo yo")

	start := time.Now()
	// Create a channel of string type each time!!!!
	ch := make(chan string)

	// They can do multiple urls at a time
	for index, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			//fmt.Fprintf(os.Stderr, "missing prefix: %v\n", url)
			//os.Exit(1)
			url = "http://" + url
		}
		// Create a new goroutine here !!!
		go fetch(url, ch, index) // start a goroutine with the newly created channel
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	os.Exit(0)
}

// Interesting, they are not using a return value.  Instead
// the code uses the channel parameter to return results.
func fetch(url string, ch chan<- string, counter int) {
	var filename string

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// ORIG ch <- fmt.Sprintf(err) // send error to channel ch
		ch <- fmt.Sprintf("error: %s\n", err.Error()) // send (formatted string method,) error to channel ch
		return
	}

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
	//fmt.Printf("== writeTextToFile(bytes, %s)\n", filename)
	//body_as_text := string(bytes)

	//fmt.Println("-------")
	//fmt.Printf("content %s: \n", body_as_text)
	//fmt.Println("-------")

	// Determine timestamp for file suffix
	// Load a location
	//loc, err := time.LoadLocation("America/New_York")
	//if err != nil {
	//	fmt.Println("Error loading location:", err)
	//	return
	//}
	// Get the current time in that location
	//currentTime := time.Now().In(loc)

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

// A web server
// * returns lissajous gif
//
// Demos how returning a binary to the client via
// another module can be done.
//
//
// ch 1.7
//
// USAGE/DEMO:
// Run this server and then use a web browser to view results
//
// Other notes
//
// all output mechanisms share a common interface via io.Writer
//
// * fetchall discards via ioutil.Discard
// * webserver writes to http.ResponseWriter
// * printf writes to os.Stdout
// add some text

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/netskink/lissajous"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	var cycles int = 5

	fmt.Println("==== handler() ====")
	fmt.Printf("Path: %s\n", r.URL.Path)
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL)
	fmt.Printf("Protocol: %s\n", r.Proto)
	for k, v := range r.Header {
		fmt.Printf("Header[%q]: %q\n", k, v)
	}
	fmt.Printf("Host: %q\n", r.Host)
	fmt.Printf("RemoteAddr: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		// https://pkg.go.dev/log
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Printf("Form[%q]: %q\n", k, v)
		if k == "cycles" {
			var err error
			cycles, err = strconv.Atoi(v[0])
			if err != nil {
				cycles = 5
			}
		}
	}

	lissajous.Lissajous(w, cycles)
}

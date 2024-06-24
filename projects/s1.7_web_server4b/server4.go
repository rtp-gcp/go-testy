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

package main

import (
	"log"
	"net/http"
)

func main() {
	// This is an anonymous function, he calls this
	// a function literal
	handler := func(w http.ResponseWriter, r *http.Request) {
		// this uses a routine also in main package, its in
		// a file in this same dir.
		lissajous(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

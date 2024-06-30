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
		// This uses a module which needs to be installed
		// cd ~..../projects/lissajous
		// go build <-- previously done via ninja build
		// go install <-- this installs the module in ~/go/bin/ workspace dir
		lissajous(w)
	}
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

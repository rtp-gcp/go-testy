// A web server ( echo server with status url )
// ch 1.7
//
// Given a url with a path component such as
//
// protocol | port     |  host | domain | path
// http://  (80 default) www.skink.net/foo.html
//
// This server would return for http://localhost:8000/hello
// URL.Path = "hello".
//
// Adds a status url.  The status returns the number of requests
// done to the server.
// http://localhost:8000/count
//
// USAGE/DEMO:
// Run this server and then use the fetchall from 1.6 first
// exercise (s1.6_fetchall_ex) to be the client.
// Remember, the handler routines must be called to increment
// the usage counter.  Calling the /count url path does
// not increment the counter
//
// Returns count
// ./fetchall http://localhost:8000/count; cat file*; rm file*
//
// Increments count
// ./fetchall http://localhost:8000/yo

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)      // each request calls handler
	http.HandleFunc("/count", counter) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echose the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==== handler() ====")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

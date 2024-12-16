// A web server
//
// echo server with :
// * status url
// * returns headers and form data
//
// It uses the log package to log errors which go to stderr.
// https://pkg.go.dev/log
//
// Besides returning, the path name to the client, it
// adds the addition of the header contents and url
// components.
//
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
//
// Test key value parms for form (Query parameters?)
// For notes on query parameters (?x=y) vs path parameters (/foo/goo)
// https://github.com/netskink/postman-testy/blob/main/notes.md
//
//./fetchall http://localhost:8000/yo\?mykey\=myval\&ak\=av; cat file*; rm file*

// Other notes
//
// all output mechanisms share a common interface via io.Writer
//
// * fetchall discards via ioutil.Discard
// * webserver writes to http.ResponseWriter
// * printf writes to os.Stdout

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
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q]: %q\n", k, v)
	}
	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		// https://pkg.go.dev/log
		log.Print(err)
	}

	// He says, he could write it like this but
	// combining statements make it shorter and reduces
	// scope of the err variable so that it only exists
	// in the if stanza/block
	// err := r.ParseForm()
	// if err != nil {
	//  	log.Print(err)
	// }

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q]: %q\n", k, v)
	}
	mu.Lock()
	count++
	mu.Unlock()
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}



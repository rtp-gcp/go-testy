// A web server ( minimal echo server )
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
// USAGE/DEMO:
// Run this server and then use the fetchall from 1.6 first
// exercise (s1.6_fetchall_ex) to be the client.

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echose the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	// in sprintf format
	// * q is a quoted string
	// * v is for value
	// * T is for type
	// * t is for boolean
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// A web server
//
// Given a url with a path component such as
//
// protocol | port     |  host | domain | path
// http://  (80 default) www.skink.net/foo.html
//
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
        "html/template"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
    dir := http.Dir("./static")
    fs := http.FileServer(dir)
    fs = http.StripPrefix("/static", fs)

    // handlers by prefix
    http.Handle("/static/", fs)      // static files needs a /static prefix in url
    http.HandleFunc("/handler", handler)      
    http.HandleFunc("/count", counter) 
    http.HandleFunc("/simpletemplate", simpletemplate) // html simpletemplate needs /simpletemplate
    http.HandleFunc("/persontemplate", persontemplate) // html persontemplate needs /persontemplate

    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func simpletemplate(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("simple.html")
    t.Execute(w, "Hello World!")
}

type Person struct {
    Name        string
    Gender      string
    Homeworld   string
}

func persontemplate(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("person.html")
    t.Execute(w, Person {
        Name: "Luke Skywalker",
        Gender: "male",
        Homeworld: "Tatooine",
    })
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



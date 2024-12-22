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
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"

	"google.golang.org/api/iterator"
)

var (
	mu                  sync.Mutex
	count               int
	ctx                 context.Context
	client              *firestore.Client
	users               *firestore.CollectionRef
	answers             *firestore.CollectionRef
	questionsCollection *firestore.CollectionRef
	userDocs            *firestore.DocumentRef
	ada                 *firestore.DocumentRef
)

func main() {
	fmt.Println("===== fs1:main() ====")

	//
	// firestore init
	//

	// from authentication example
	// client, err := storage.NewClient(ctx, option.WithCredentialsFile("../../../keys/shawdavis-244405e0b7c4.json"))
	ctx = context.Background()

	// Fetch Credentials from ENV variable
	credentialsFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentialsFile == "" {
		log.Fatal("Environment variable for credentials path not found.")
	}

	// from api on firestore
	client, err := firestore.NewClient(ctx, "shawdavis", option.WithCredentialsFile(credentialsFile))
	if err != nil {
		fmt.Println("error opening client to project")
		log.Fatal(err) // log.Print() will simply print. log.Fatal() will print and exit.
	}
	// Just to use the client in a print
	fmt.Printf("client: %v\n", client)
	fmt.Printf("Type of variable for client: %T\n", client)

	//
	// READ - this is just to verify we can read from server. If it
	// fails, it will abort before we start webserver.
	//

	// Attempt to read the database
	users = client.Collection("users")
	questionsCollection = client.Collection("questions")
	answers = client.Collection("answers")
	// fmt.Printf("Type of variable for users: %T\n", users)
	ada = users.Doc("alovelace") // This gets a documentRef
	// fmt.Printf("Type of variable for ada: %T\n", ada)
	// the actual read - this does network traffic
	docsnap, err := ada.Get(ctx)
	if err != nil {
		fmt.Println("error getting data")
		log.Fatal(err)
	}
	dataMap := docsnap.Data()
	fmt.Println(dataMap)

	//
	// Webserver init
	//

	dir := http.Dir("./static")
	fs := http.FileServer(dir)
	fs = http.StripPrefix("/static", fs)

	// handlers by prefix
	http.Handle("/static/", fs)     // static files needs a /static prefix in url
	http.HandleFunc("/", questions) // main entry page
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/simpletemplate", simpletemplate) // html simpletemplate needs /simpletemplate
	http.HandleFunc("/persontemplate", persontemplate) // html persontemplate needs /persontemplate

	// this is a handler to show the page and a handler to process the response
	http.HandleFunc("/form", form)
	http.HandleFunc("/process", process)

	// Where we start to talk to firestore
	http.HandleFunc("/userform", userform)
	http.HandleFunc("/userresponse", userresponse)
	// This routine does not exit unless and error occurs
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

/////////////vvvvvvvvvvvvvvv/////////////

//type Question struct {
//        ID string           `firestore:"ID"`
//        answers string       `firestore:"answers"`
//        question_text string        `firestore:"question_text"`
//}

type Questions struct {
	userDocReference *firestore.DocumentRef
	question         string
}

type Questions2 struct {
	UserID   string
	Question string
}

func questions(w http.ResponseWriter, r *http.Request) {
	var questionsList [10]Questions
	var questions2List [2]Questions2
	i := 0

	iter := questionsCollection.Documents(ctx)

	for {
		docsnap, err := iter.Next()
		if err == iterator.Done {
			break // no more documents
		}
		if err != nil {
			log.Fatalf("Failed to fetch documents: %v", err)
		}

		// Print the document ID and optionally the fields
		// fmt.Printf("Document ID: %s\n", docsnap.Ref.ID)
		// fmt.Printf("Fields: %v\n", docsnap.Data())

		// fmt.Printf("Type of variable for docsnap: %T\n", docsnap)

		// add to array for output
		refe := docsnap.Data()["user"].(*firestore.DocumentRef)
		// fmt.Println(refe)
		// fmt.Printf("Type of variable for refe: %T\n", refe)
		questionsList[i].userDocReference = refe

		quest := docsnap.Data()["question_text"].(string)
		questionsList[i].question = quest
		i++
	}

	for j := 0; j < i; j++ {
		//fmt.Println(questionsList[j].userDocReference)
		//fmt.Println(questionsList[j].userDocReference.ID)
		//fmt.Println(questionsList[j].question)
		//
		questions2List[j].UserID = questionsList[j].userDocReference.ID
		questions2List[j].Question = questionsList[j].question
	}

	for j := 0; j < i; j++ {
		fmt.Println(questions2List[j].UserID)
		fmt.Println(questions2List[j].Question)
	}

	t, _ := template.ParseFiles("questionsform.html")
	err := t.Execute(w, questions2List)
	// foo := []Questions2{{UserID: "123", Question: "xxxx"},{UserID: "456", Question: "yyyy"}}
	// err := t.Execute(w,foo)
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
	}

	mu.Lock()
	count++
	mu.Unlock()
}

/////////////^^^^^^^^^^^^^^^^////////////

// ///////////vvvvvvvvvvvvvvv/////////////
func userform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("userform.html")
	err := t.Execute(w, "")
	if err != nil {
		fmt.Println("An error occurred")
		fmt.Println(err)
	}
	mu.Lock()
	count++
	mu.Unlock()
}

type User struct {
	Year      float64 `firestore:"born"`
	FirstName string  `firestore:"first"`
	LastName  string  `firestore:"last"`
}

func userresponse(w http.ResponseWriter, r *http.Request) {
	var fName string                         // will hold the first name user specified on form
	var querydoc *firestore.DocumentSnapshot // holds result of query

	r.ParseForm()

	// We can print the values back to the webpage
	// which is now /process
	for k, v := range r.Form {
		fmt.Printf("%s: %s\n", k, v)
	}

	// We know which name was specified in r.Form[key=FirstName] with value=[user input]
	if values, ok := r.Form["fname"]; ok && len(values) > 0 {
		fName = values[0]
		fmt.Printf("== userresponse() fname: %s\n", fName)
	} else {
		fmt.Println("== userresponse() fname: fname is not provided")
	}

	// if the fname is not specified on the form, nothing to query, just print an error
	// message or return early

	if fName == "" {
		fmt.Println("== userresponse() early return")
		return
	}

	// Query firestore to get info for the specified user
	// TODO: this iterates all docs that match, we only print
	// the last result.
	q := users.Where("first", "==", fName)
	// use query
	iter := q.Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		fmt.Printf("Type of variable for doc: %T\n", doc)
		if err == iterator.Done {
			break
		}
		if err != nil {
			// do something if an error
			fmt.Println("error while iterating firestore results")
			log.Fatal(err)
		}
		fmt.Println(doc.Data())
		querydoc = doc
	}
	fmt.Println(querydoc.Data())

	var userData User
	if err := querydoc.DataTo(&userData); err != nil {
		fmt.Println("error extracting data from doc snap")
		log.Fatal(err)
	}
	fmt.Printf("Year: %v\n", userData.Year)
	fmt.Printf("FirstName: %v\n", userData.FirstName)
	fmt.Printf("LastName: %v\n", userData.LastName)

	// update the webpage with data for the selected user.

	//
	t, _ := template.ParseFiles("userresponse.html")
	t.Execute(w, User{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Year:      userData.Year,
	})
}

/////////////^^^^^^^^^^^^^^^^////////////

// ///////////vvvvvvvvvvvvvvv/////////////
func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, "sample form")
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// We can print the values back to the webpage
	// which is now /process
	for k, v := range r.Form {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}

	//
	//t, _ := template.ParseFiles("simple.html")
	//t.Execute(w, "Hello World!")
}

/////////////^^^^^^^^^^^^^^^^////////////

func simpletemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("simple.html")
	t.Execute(w, "Hello World!")
}

type Person struct {
	Name      string
	Gender    string
	Homeworld string
}

func persontemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("person.html")
	t.Execute(w, Person{
		Name:      "Luke Skywalker",
		Gender:    "male",
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

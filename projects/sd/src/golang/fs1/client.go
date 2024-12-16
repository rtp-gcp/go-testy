package main

import (
    //"cloud.google.com/go"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/api/iterator"
    "log"
    "fmt"
    "context"
)


func main() {
    fmt.Println("===== fs1:main() ====")

    // from authentication example
    //client, err := storage.NewClient(ctx, option.WithCredentialsFile("../../../keys/shawdavis-244405e0b7c4.json"))
    ctx := context.Background()
    
    // from api on firestore
    client, err := firestore.NewClient(ctx, "shawdavis", option.WithCredentialsFile("../../../keys/shawdavis-244405e0b7c4.json"))
    if err != nil {
        fmt.Println("error opening client to project")
        log.Fatal(err)  // log.Print() will simply print. log.Fatal() will print and exit.
    }
    // Just to use the client in a print
    fmt.Printf("client: %v\n", client)

    //
    // READ
    //

    // Attempt to read the database
    users := client.Collection("users")
    ada := users.Doc("alovelace")  // This gets a documentRef
    // the actual read - this does network traffic
    docsnap, err := ada.Get(ctx)
    if err != nil {
        fmt.Println("error getting data")
        log.Fatal(err)
    }
    dataMap := docsnap.Data()
    fmt.Println(dataMap)


    // Or, in a single call:
    ada = client.Doc("users/alovelace")
    // the actual read - this does network traffic
    docsnap, err = ada.Get(ctx)
    if err != nil {
        fmt.Println("error getting data")
        log.Fatal(err)
    }
    dataMap = docsnap.Data()
    fmt.Println(dataMap)

    // You can also obtain a single field with DataAt or extract data into a struct with DataTo using
    // a type definition
    type User struct {
        Year float64           `firestore:"born"`
        FirstName string       `firestore:"first"`
        LastName string        `firestore:"last"`
    }

    var adaData User
    if err = docsnap.DataTo(&adaData); err != nil {
        fmt.Println("error extracting data from docsnap")
        log.Fatal(err)
    }
    fmt.Printf("Year: %v\n", adaData.Year)
    fmt.Printf("FirstName: %v\n", adaData.FirstName)
    fmt.Printf("LastName: %v\n", adaData.LastName)


    //
    // WRITE
    //

    // write a new individual document to a collection
    russell := users.Doc("rshaw")  // This gets a documentRef
    wr, err := russell.Create(ctx, User {  // Create does a new doc. Set creates/replaces and existing doc.
        Year: 1969,
        FirstName: "Russell",
        LastName: "Shaw",
    })
    if err != nil {
        fmt.Println("error writing data ")
        if status.Code(err) == codes.AlreadyExists {
            fmt.Print("already exists. no problem. continue.\n");
        } else {
            log.Fatal(err)
        }
    } else {
        fmt.Println(wr)
    }


    //
    // Multiple Reads
    //
    fmt.Println("**** multiple reads example ***")
    docsnaps, err := client.GetAll(ctx, []*firestore.DocumentRef{
        users.Doc("alovelace"), users.Doc("rshaw"),
    })
    if err != nil {
        log.Fatal(err)
    }

    var userData User
    for _, ds := range docsnaps {
        if err = ds.DataTo(&userData); err != nil {
            fmt.Println("error extracting data from docsnap")
            log.Fatal(err)
        }
        fmt.Printf("Year: %v\n", userData.Year)
        fmt.Printf("FirstName: %v\n", userData.FirstName)
        fmt.Printf("LastName: %v\n", userData.LastName)
    }
    fmt.Println("**** multiple reads example end ***")

    //
    // UPDATE
    //

    // Change Ada to have a birt year of 1915
    wr, err = ada.Update(ctx, []firestore.Update{
        {Path: "born", Value: 1915},
    })
    if err != nil {
    }
    fmt.Println(wr.UpdateTime)

    //
    // QUERY
    //

    fmt.Println("**** query example ***")
    // specify query
    q := users.Where("born", ">", 1900).OrderBy("born", firestore.Desc)
    // use query
    iter := q.Documents(ctx)
    defer iter.Stop()
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
        }
        fmt.Println(doc.Data())
    }
    fmt.Println("**** query example end ***")


    // Change Ada back so it has a birth year of 1815
    wr, err = ada.Update(ctx, []firestore.Update{
        {Path: "born", Value: 1815},
    })
    if err != nil {
    }
    fmt.Println(wr.UpdateTime)

    //
    // QUERY 2
    //

    fmt.Println("**** query 2 example ***")
    // use query
    iter = q.Documents(ctx)  // This causes network usage
    defer iter.Stop()
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
        }
        fmt.Println(doc.Data())
    }
    fmt.Println("**** query 2 example end ***")



    // 
    // END
    //
    fmt.Println("Normal exit.")
}



# golang code

The hackathon attempts to build a webserver in go/golang

# Code Samples

* ws1
    - a webserver in go
* fs1
    - a client for gcp firestore in go


# workflow for go ws1

This is the web server sample. It does not depend on any external modules.

## macOS

Build and run from scratch

1. go mod init server
2. go mod tidy
3. go build server
4. ./server

Build and run from existing code

1. go build server
2. ./server


# workflow for go fs1

This is the client for the gcp firestore. It depends upon the google cloud module.
The reference for the [module](https://pkg.go.dev/cloud.google.com/go#section-readme)
specifies that you need to do a `go get` command. eg. :

```
go get cloud.google.com/go/firestore # Replace with the package you want to use.
```

However, the go mod tidy appears to do this already.


## macOS

Build and run from scratch

1. go mod init client
2. go mod tidy
3. go build client
4. ./client

Build and run from existing code

1. go build client
2. ./client


## URLS

* https://cloud.google.com/go/docs/reference



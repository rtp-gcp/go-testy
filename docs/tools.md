# tool chain commands


## go build

in root dir

```
go build -o app
./app
```

The demo github has the test app icon on left side.

## go env

```
go env
```

Important variables are 

* GOROOT defines where your Go SDK has been installed to, and 
* GOPATH defines the root of your workspace, and will be the location where a lot of 3rd party source code, packages, and/or modules will be found
* GOENV is where your env is
* `go help environment` will list info on each variable

Settings from `go env` output:

```
GOENV="/home/davis/.config/go/env"
GOROOT="/usr/lib/go-1.15"
GOPATH="/home/davis/go"
```

Display help on each variable

```
go help env
```

## go doc

Brings up documentation for a package.  Consider it the equvialent of a manpage.

```
go doc
go doc math
go doc math.Pow10
```

## go run

Build and run in one command

```
go run main.go
```



## go build

can build cross platform

```
go build
go build -o testy
GOOS=linux GOARCH=amd64 go build -o godemo-linux-amd64 .
GOOS=windows GOARCH=amd64 go build -o godemo-win-amd64 .
go tool dist list
```

## go test

```
go test .
# test two projects simultaneously
go test . util
go test -v -run=^TestGetX$ ./...
```

## go get

Get a module of a specific version

```
go get github.com/google/uuid@1.1.0
# get a specific commit 
go get github.com/google/uuid@d460ce9
cat go.mod
```



## go list

see all of the dependencies that the current project is configured with. This will also now be reflected with a go.mod file: 

```
go list -m all
```

List standard libraries

```
go list std
```



## go mod

```
go mod init
rm go.mod go.sum
go mod init github.com/cloudacadmy/godemo
ls -la $(go env GOPATH)/pkg/mod/github.com/google

go mod verify
```

## go get

```
go get
cat go.mod
```

## go list

list out all dependent modules and whether an updated version exists and if so, what the latest semantic version number is.

```
go list -m all
go list -u -m all", used to 
```

## go mod

Prune out any required module declarations from the go.mod file if the module dependency is no longer used within the project source code anywhere.

```
go mod why -m github.com/google/uuid
# examine md5sum
go mod verify
go mod tidy  
```
## go list

```
go list -u -m all
```

## go fmt

Used to reformat source code to what Go considers best practices. This 


```
go fmt 
```

 This will reformat the code showing just the differences but not save the changes. The -s flag performs simplifications on the code. The -d flag performs a diff to highlight before and after. The absence of the -w flag allows you to examine the changes before applying them to the disk. Extending the same command with the -w flag will cause the changes to be saved. 

```
# preview changes
gofmt -s -d main.go
# write the changes
gofmt -s -d -w main.go
# write the changes to all files recursively
gofmt -s -d -w .
# alternativewrite the changes to all files recursively
go fmt  ./...
```


To perform this command recursively, use the following command: The go fmt ./... is a shortcut wrapper over the gofmt tool, and will perform the same reformatting recursively across the entire project, saving all updates to the disk in one hit. 

```
gofmt -s -d -w . 
go fmt ./..
```


## go vet

The "go vet" command is used to perform a *static* *analysis* across your source code. It can be used to detect and report on code that although compiles may be considered problematic. For example, it can detect issues such as unreachable code. go vet can be called on single files, all files in the current directory, or across the entire project. As an example, if I were to add the following bad function implementation to the util calc.go file and then compiled it, it should compile even though it contains unreachable code, as it does. Let's now run go vet on it, and as you can see the static analysis has flagged the unreachable code. 

```
go vet ./...
```

## golin

TODO

## go imports

* goimports
  - inserts/removes packages from the import section.
  - install
    * `go install golang.org/x/tools/cmd/goimports@latest`
    * This install to $GOPATH/bin.  Note: `$ go env` will show $GOPATH if not in ENV.  if this dir is not 
    in $PATH, simply add it.
  - usage
    * `goimports -w main.go` will rewrite main.go source file

## go format


* gofmt
  - formats code in canonical form.
  - `gofmt -w main.go` will rewrite the main.go source file


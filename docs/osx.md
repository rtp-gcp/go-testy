# notes on go on osx

## Where is your workspace?

```
go env GOPATH
/Users/davis/go
cd go
mkdir src
cd src
mkdir example
cd example
vi hello.go
go run hello.go
```


# build

This did not build for me as is, but once I read this website 
I did the module thing and then `go build` would work.

[create-module](https://golang.org/doc/tutorial/create-modulex)


```
go mod init skink.net/hello
go build 
```

This built a `go.mod` file and a `hello` executable.

# install

This will create a exe in `~/go/bin/hello`

```
go install
```




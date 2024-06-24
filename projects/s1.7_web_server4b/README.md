# section 1.7 server

This is one package implementation:

* server4.go and lissajous.go are both in main package
* lissajous.go modified to remove main() call.
* lissajous.go entry point change from capital to lowercase since no need for a public function since its same package

build.ninja works, but here is a working alternative for cmdline

```
go mod init st1.7_web_server4b
go mod tidy
go vet
go build 
./s17_web_server4b
or
go run .
```

open browser to `http://localhost:8000`

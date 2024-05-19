
```
cd ..../go-testy/genai
go get gopl.io/ch1/helloword
$GOPATH/bin/helloworld
```

Odd, it says needs golang 1.5 but latest is 1.22 and I have that installed.


```
go run helloworld.go
```
## Workspace

The setenv.sh will set the gopath to the root of this folder.  It makes a workspace.  Each project, or separate program
needs to be in a subdir of this workspace.

### setup

After this all the go build and go run code will work.

#### do this once
```
$ go mod init myhellorld
```

#### do this afterwards or each time the imports changes

```
$ go mod tidy
```

#### do this to build

````
$ go build myhelloworld
```
This makes a file named myhelloword which can be run.  


## install goimports 

goimports manages the insertion and removaol of import declarations as needed. Its not part of the
standard distribution.





#### GOPATH VS GOROOT
In Go programming language, both GOPATH and GOROOT are environment variables used to define important paths for the Go toolchain and development workflow, but they serve different purposes:

1. **GOPATH**:
   - GOPATH is an environment variable that specifies the root of your workspace. It tells Go where to look for your Go code and where to place compiled binaries.
   - Inside the GOPATH, there are three main directories: `src`, `bin`, and `pkg`. The `src` directory is where your Go source code resides, `bin` is where compiled binaries are stored after you run `go install`, and `pkg` contains package objects.
   - You can have multiple workspaces by setting different GOPATH values.

2. **GOROOT**:
   - GOROOT is an environment variable that points to the location where Go is installed on your system. It is set during the installation of the Go programming language.
   - GOROOT tells the Go toolchain where to find system-wide Go-related files, such as standard libraries, compiler tools, and other essential components.
   - It's generally recommended not to modify GOROOT after installation unless you are updating or reinstalling Go.

In summary, GOPATH is where your own Go projects reside, while GOROOT points to the Go installation directory, including the standard library and Go tools.


# URLS

* [code settings](https://github.com/golang/vscode-go/blob/master/docs/settings.md)
* [golang book online](https://www.informit.com/articles/article.aspx?p=2453564&seqNum=4)



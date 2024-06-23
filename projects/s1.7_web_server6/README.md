# section 1.7 server

A package in a complete differt root dir and in github, 
and we import from github implementation.

This uses the lissajous in github aka ~/progs/lissajous

Make sure you don't have these files around before you do
do the go mod init.
* go.sum
* go.mod
* ~/davis/go/pkg/mod/cache/download/github.com/netskink/lissajous

TODO: We had some issue where it was not using the latest version in github.  Our original code had a bug, we fixed it
and each time it was doing mod tidy, it was downloading the
old version.  When we finally got it to work, we deleted the
go.sum file and the package in ~/go/pk/mod/cache.

TODO: Need to figure out ow to do this via go commands.


```
go mod init st1.7_web_server6
```
This will add the lisajous module from github

```
go mod tidy
```


```
go vet
go run .
```



open browser to `http://localhost:8000`

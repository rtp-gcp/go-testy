# section 1.7 server

A package in a complete differt root dir and in github, 
and we import from github implementation.

This uses the lissajous in github aka ~/progs/lissajous

Make sure you don't have these files around before you do
do the go mod init.
* go.sum
* go.mod
* had to perform `go clean -modcache`
    * This removes ~/go/pkg/mod/cache/download/github.com/netskink/lissajous
    * This removes ~/go/pkg/mod/github.com/netskink/lissajous@000..n-hash
* We also deleted a lissajous file in ~/go/bin/lissajous.

TODO: We had some issue where it was not using the latest version in github.  Our original code had a bug, we fixed it
and each time it was doing mod tidy, it was downloading the
old version.  When we finally got it to work, we deleted the
go.sum file and the package in ~/go/pkg/mod/github.  Deleting the one in cache/download
did not resolve the problem.  Once, we also did the `go clean -modcache` command it would then pull the latest code from github.  However, afterwards, we made additional mods and despite repeating this workflow, it continued to 
use the same code and did not pull from github as it did when it worked.  OMG, once we deleted the bin file 
entry, then it would fetch the latest source from git.

However, we attempted to do this once more and force an error and it did not fetch from git the latest code.?/?!!
The file that was ~/go/bin/lissajous was not recreated as a result of building this source was not recreated.
Perhaps it was put there a long time ago as part of a `go install .`?

TODO: Punt for now and figure out versions.


TODO: Need to figure out ow to do this via go commands.

This shows the current env settings used by go
```
$ go env
```

This cleans the ~/go/pkg/ folder contents.

```
$ go clean -modcache
```

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

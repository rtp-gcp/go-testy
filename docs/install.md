# install and base commands

```
$ which go
/usr/bin/go
davis@zatoichi:~/progs/go-testy 
$ ls -l /usr/bin/go
lrwxrwxrwx 1 root root 21 Sep 16  2020 /usr/bin/go -> ../lib/go-1.15/bin/go
```


## go version

## go env

* GOROOT where go sdk is located.  includes std library, toolchain. 
* GOPATH where your workspace is
* GOENV is where your env is

```
GOENV="/home/davis/.config/go/env"
GOROOT="/usr/lib/go-1.15"
GOPATH="/home/davis/go"
```

Put your src in `/home/davis/go/src` and make a link to this repo.  Alternatively there is a `go modules` thing.

## go run

compile and run a go program.

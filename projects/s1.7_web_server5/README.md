# section 1.7 server

A package in a complete differt root dir, and we
import from local packages implementation.

This uses the lissajous2 in this repo .

```
go mod init st1.7_web_server5
```

Had to add this line to the go.mod file

```
replace lissajous2 => /Users/davis/progs/go-testy/projects/lissajous2
```

```
go mod tidy
```

This added the require line.  When we did it before it
was looking for the module in ~/go/src which was not there.

```
replace lissajous2 => /Users/davis/progs/go-testy/projects/lissajous2
```

```
go vet
go run .
```



open browser to `http://localhost:8000`

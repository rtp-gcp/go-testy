# notes


# tests

```
$ cd ex2.3
$ og mod tidy  --- to pull in the testing framework and the popcount package
$ go test -v ./...    - to recursively run all tests in subdirectories
```

## to run a specific test

```
$ cd popcount
$ go test -v -run ^TestFunction64k$
```

## to benchmark a test

Need to add some test code to call a function multiple times:

```
func BenchmarkFunction(b *testing.B) {
        for i := 0; i < b.N; i++ {
                popcount.PopCount(64 * 1024)
        }
}
```

This will run a specific test and then call the benchmark test
```
$ cd popcount
$ go test -v -run ^TestFunction64k$ -bench=BenchmarkFunction
```

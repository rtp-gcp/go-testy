# section 1.6 Fetching URLs concurrently

One of the capabilities of golang is the built-in ability to perform multiple
tasks simultaneously.

This example introduces Golangs concurrency mechanisms:

* goroutines
* channels

This example fetches multiple URLs simulatenously.  The beauty is that all
fetches only take as much time to fetch as the longest particular fetch to occur.
This example discards the responses/results but reports the size and time to fetch.

# Getting help with go command line

```
$ go doc http.Get
```

# Notes on the code example

* A goroutine is a concurrent function execution.
* A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. 
    - The function main runs in a goroutine and the go statement creates additional goroutines.
* When one goroutine attempts to send/receive on a channel, it blocks until another
goroutine does the appropriate method.  ie. one blocks on read until the other performs write.
    - in this example, fetch writes a string buffer and main reads the response.
        - fetch()
            * `ch <- fmt.Sprintf("%.2fs  %7d %s", secs, nbytes, url)`
        - main()
            * `fmt.Println(<-ch) // receive from channel ch`
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

# Exercises

## Exercise 1.10

Modify the code to save the results with a file name including the date and timestamp.  
Investigate caching by running fetchall twice in succession on a large website.  Record the
reported time changes.  

Ask yourself:

* Do the times vary?
    - yes.  Google first time 30s, second time 17s.  Kind of the same thing for amazon as well.
* Does the returned results written to file vary?
    - yes.  For google, the nonce values change.


## Exercise 1.11

Try fetchall with longer argument lists such as samples from the top million web sites available at
alexa.com.  How does the program respond if a web site does not respond?  (Section 8.9 describes 
mechanisms for coping in such cases.)

This domain is gone, but google points to [wikipedia](https://en.wikipedia.org/wiki/List_of_most-visited_websites)
    1. google
    2. youtube
    3. facebook
    4. instagram
    5. x/twitter





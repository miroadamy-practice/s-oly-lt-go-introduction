package main

// NON-CONC

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

)

// responseSize retrieves "url" and prints
// the response length in bytes.

func responseSize3(url string, channel chan int) {
	fmt.Println("Getting", url)
	// Note: errors ignored with _!
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	channel <- len(body)
}

func main() {
	start := time.Now() // Unchanged
	// Make a channel to carry ints.
	sizes := make(chan int)
	// Pass channel to each call to responseSize.
	go responseSize3("https://example.com/", sizes)
	go responseSize3("https://golang.org/", sizes)
	go responseSize3("https://golang.org/doc", sizes)
	// Read and print values from channel.
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(time.Since(start).Seconds()) // Unchanged
}


/*
Getting https://golang.org/doc
Getting https://golang.org/
Getting https://example.com/
1270
8158
12558
0.695384291


---

Finishes in half the time of the original! (YMMV.)
• The channel accomplishes two things:
• Channel reads cause main goroutine to block until responseSize goroutines send, so they have
time to finish before program ends.
• The channel transmits data from the responseSize goroutines back to the main goroutine.
*/

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
func responseSize(url string) {
	fmt.Println("Getting", url)
	// Note: errors ignored with _!
	response, _ := http.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(len(body))
}

func main() {
	// Note the time we started.
	start := time.Now()
	responseSize("https://example.com/")
	responseSize("https://golang.org/")
	responseSize("https://golang.org/doc")
	// Print how long everything took.
	fmt.Println(time.Since(start).Seconds(), "seconds")
}

/*
Getting https://example.com/
1270
Getting https://golang.org/
8158
Getting https://golang.org/doc
12558
1.5341211000000001 seconds
*/

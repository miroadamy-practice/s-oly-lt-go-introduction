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
// UNCHANGED func responseSize(url string) {
func responseSize2(url string) {
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
	go responseSize2("https://example.com/")
	go responseSize2("https://golang.org/")
	go responseSize2("https://golang.org/doc")
	// Print how long everything took.
	fmt.Println(time.Since(start).Seconds(), "seconds")
}

/*
Getting https://example.com/
Getting https://golang.org/
9.378e-06 seconds
Getting https://golang.org/doc

DOES NOT WAIT to finish
*/

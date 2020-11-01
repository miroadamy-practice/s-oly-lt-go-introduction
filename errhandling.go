package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// It’s usually polite to end conversations with “goodbye”:
func Socialize() {

	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
	fmt.Println("Goodbye!")
}

func Socialize2() {
	// This call will be made when Socialize ends.
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}


func Socialize3() error {
	// Deferred call is made even if Socialize
	// exits early (say, due to an error).
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk.")
	// The below code won't be run!
	fmt.Println("Nice weather, eh?")
	return nil
}

func PrintLines(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func main() {
	Socialize()
	Socialize2()
	err := Socialize3()
	if err != nil {
		log.Fatal(err)
	}

	// more realistic example
	err2 := PrintLines("lorem_ipsum.txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	Socialize4()
}


// panic usually signals an unanticipated error.
// This example is just to show its mechanics.

func Socialize4() {
	fmt.Println("Hello!")
	panic("I need to get out of here!")
	// The below code won't be run!
	fmt.Println("Nice weather, eh?")
	fmt.Println("Goodbye!")
}

/*
Hello!
panic: I need to get out of here.
goroutine 1 [running]:
main.Socialize()
 /Users/jay/socialize4_panic.go:9 +0x79
main.main()
 /Users/jay/socialize4_panic.go:16 +0x20
exit status 2
 */

func Socialize5() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	panic("I need to get out of here!")
	// The below code won't be run!
	fmt.Println("Nice weather, eh?")
}

/*

Hello!
Goodbye!
panic: I need to get out of here!
goroutine 1 [running]:
main.Socialize()
 /Users/jay/socialize5_panic_defer.go:10 +0xd5
main.main()
 /Users/jay/socialize5_panic_defer.go:16 +0x20
exit status 2

 */


func CalmDown() {
	// Halt the panic.
	panicValue := recover()
	// Print value passed to panic().
	fmt.Println(panicValue)
}
func Socialize6() {
	defer fmt.Println("Goodbye!")
	defer CalmDown()
	fmt.Println("Hello!")
	panic("I need to get out of here!")
	// The below code won't be run!
	fmt.Println("Nice weather, eh?")
}

/*
Hello!
I need to get out of here!
Goodbye!

 */

// “panic” should not be used like an exception
//I know of one place in the standard library that panic is used in normal program flow:
//in a recursive parsing function that panics to unwind the call stack after a parsing
//error. (The function then recovers and handles the error normally.)


// Generally, panic should be used only to indicate “impossible” situations:
func awardPrize() {
	doorNumber := rand.Intn(3) + 1
	if doorNumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {
		fmt.Println("You win a car!")
	} else if doorNumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		// This should never happen.
		panic("invalid door number")
	}
}

// Google “golang errors are values” (which should take you to https://blog.golang.org/
// errors-are-values) for some tips on making error handling more pleasant.
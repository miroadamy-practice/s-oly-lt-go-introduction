package main

import (
	"fmt"
	"greeting"
	"greeting/deutsch"
	"greeting/espanol"
	"log"
	"reflect"
	"strconv"
)

func sayHi() {
	fmt.Println("Hi!")
}

func say(phrase string, times int) {
	for i := 0; i < times; i++ {
		fmt.Print(phrase)
	}
	fmt.Print("\n")
}

func my1() {
	flag, err := strconv.ParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(flag) // => true
	flag, err = strconv.ParseBool("foobar")
	if err != nil {
		log.Fatal(err)
		// => 2009/11/10 23:00:00 strconv.ParseBool:
		// => parsing "foobar": invalid syntax
	}
	fmt.Println(flag)
}

func my2() {
	flag, err := myParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(flag) // => true
	flag, err = myParseBool("foobar")
	if err != nil {
		log.Fatal(err)
		// => 2009/11/10 23:00:00 strconv.ParseBool:
		// => parsing "foobar": invalid syntax
	}
	fmt.Println(flag)
}

func myParseBool(myString string) (bool, error) {
	if myString == "true" {
		return true, nil
	} else if myString == "false" {
		return false, nil
	} else {
		return false, fmt.Errorf("bad string %s", myString)
	}
}

func divide(one float64, two float64) (float64, error) {
	if two == 0 {
		return 0, fmt.Errorf("Cannot divide by Zero")
	}
	return one / two, nil
}

// Accept pointer instead of value
func double(number *int) {
	// Update value at pointer
	*number *= 2
}

func main() {
	fmt.Printf("hello world")
	//

	var myInteger int
	myInteger = 1
	var myFloat float64
	myFloat = 3.1415
	fmt.Println(myInteger) // => 1
	fmt.Println(myFloat) // => 3.1415
	fmt.Println(reflect.TypeOf(myInteger)) // => int
	fmt.Println(reflect.TypeOf(myFloat)) // => float64

	//conversions
	var length float64 = 1.2
	var width int = 2
	// But you can if you do type conversions!
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))


	// IF
	if 1 < 2 {
		fmt.Println("It's true!")
	}

	// for
	for x := 4; x <= 6; x++ {
		fmt.Println("x is now", x)
	}

	// call finctoon
	sayHi()

	say(
		"Hi",
		4,
	) // => HiHiHiHi
	say("Bye", 2) // => ByeBye

	// These will end the program
	//my1()
	//my2()

	quotient, err := divide(5.6, 0)
	if err == nil {
		fmt.Printf("%0.2f\n", quotient) // => 2.80
	} else {
		fmt.Println(err)
	}
	// POINTER
	fmt.Println(&myInteger) // => 0x1040a128
	fmt.Println(&myFloat) // => 0x1040a140
	var myBool bool
	fmt.Println(&myBool) // => 0x1040a148

	fmt.Println(reflect.TypeOf(&myInteger)) // => *int
	fmt.Println(reflect.TypeOf(&myFloat)) // => *float64
	fmt.Println(reflect.TypeOf(&myBool)) // => *bool

	var myIntPointer *int
	myIntPointer = &myInteger
	fmt.Println(myIntPointer) // => 0x1040a128
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer) // => 0x1040a140

	myFloat = 1.2345
	myInteger = 42

	fmt.Println(myIntPointer) // => 0x1040a124
	fmt.Println(*myIntPointer) // => 42

	fmt.Println(myFloatPointer) // => 0x1040a140
	fmt.Println(*myFloatPointer) // => 1.2345

	*myIntPointer = 8
	fmt.Println(*myIntPointer) // => 8
	fmt.Println(myInteger) // => 8

	amount := 6
	// Pass pointer instead of value
	double(&amount)
	fmt.Println(amount) // => 12

	// Packages

	greeting.Hello()
	greeting.Hi()

	// try to call unexported:
	// greeting.hi()
	/*

	# _/Users/miro/prj/s-oly-lt-go-introduction
	./learning1.go:166:2: cannot refer to unexported name greeting.hi
	./learning1.go:166:2: undefined: greeting.hi

	Compilation finished with exit code 2


	 */

	deutsch.Hello()
	espanol.Hello()
}

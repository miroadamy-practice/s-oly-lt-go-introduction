package main

import "fmt"

type Liters float64
type Gallons float64

type MyType string

// Specify a "receiver parameter" within a function
// definition to make it a method. The receiver
// parameter's type will be the type the method
// gets defined on.
func (m MyType) sayHi() {
	fmt.Println("Hi")
}

func (m MyType) sayHi2() {
	fmt.Println("Hi from", m)
}

func one() {
	var carFuel Gallons
	var busFuel Liters
	// Defining a type defines a conversion
	// from the underlying type to the new type
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel) // => 10
	fmt.Println(busFuel) // => 240

	// call sayHi

	value := MyType("a MyType value")
	value.sayHi() // => Hi
	anotherValue := MyType("another value")
	anotherValue.sayHi() // => Hi

	value.sayHi2()
	anotherValue.sayHi2()
}

type MyType2 MyType

func two() {
	// Underlying type is not a superclass
	// The underlying type specifies how a type’s data will be stored, so this is OK…
	// "inherited" conversion works
	value2 := MyType2("a MyType2 value")
	fmt.Println(value2)

	// functions are NOT inherited
	// ./types2.go:53:8: value2.sayHi undefined (type MyType2 has no field or method sayHi)
	// value2.sayHi()
}

/*
Underlying type is not a superclass
“Although Go has types and methods and allows an object-oriented style of
programming, there is no type hierarchy.”
—https://golang.org/doc/faq
• There is no method inheritance!
• But there’s another way to get the same benefits…
 */
func main() {
	one()
	two()
}

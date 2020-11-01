package main

import "fmt"

type Coordinates2 struct {
	Latitude  float64
	Longitude float64
}

type Landmark2 struct {
	Name string
	// An "anonymous field"
	// Has no name of its own, just a type
	Coordinates2
}

// Methods for an embedded type get promoted too!
func (c Coordinates2) Location() string {
	return fmt.Sprintf("(%0.2f, %0.2f)", c.Latitude, c.Longitude)
}

func main() {
	var l Landmark2
	l.Name = "The Googleplex"
	// Fields for "embedded struct" are "promoted"
	l.Latitude = 37.42
	l.Longitude = -122.08
	fmt.Println(l.Name, l.Latitude, l.Longitude)
	// => The Googleplex 37.42 -122.08

	// Methods from embedded type are
	// promoted to outer type
	fmt.Println(l.Location())

}
/* Embed additional types to gain additional methods.
• You’ve heard “favor composition over inheritance”…
• Go implements that principle at the language level.

 */
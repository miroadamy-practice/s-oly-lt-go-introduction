package main

import "fmt"

func shortDeclarations() {

	fmt.Println("\nShort declarations\n")

	var myArray [3]string
	mySlice := make([]string, 2)
	myMap := make(map[string]int)
	myArray[0] = "Amy2"
	fmt.Println(myArray[0]) // => Amy
	mySlice[1] = "Jose2"
	fmt.Println(mySlice[1]) // => Jose
	myMap["Ben"] = 79
	fmt.Println(myMap["Ben"]) // => 78
}

func valuesCanBeOfAnyType() {
	fmt.Println("\nvaluesCanBeOfAnyType \n")

	var flags [2]bool // Boolean array
	flags[1] = true
	fractions := make([]float64, 3) // Float slice
	fractions[0] = 0.25
	elements := make(map[string]int) // Integer map
	elements["He"] = 2
	fmt.Println(flags[1]) // => true
	fmt.Println(fractions[0]) // => 0.25
	fmt.Println(elements["He"]) // => 2

	// maps

	binaryBooleans := make(map[bool]int)
	binaryBooleans[true] = 1
	binaryBooleans[false] = 0
	fmt.Println(binaryBooleans[false]) // => 0
	fmt.Println(binaryBooleans[true]) // => 1
}
func expandingSlices() {
	fmt.Println("\nExpanding slices\n")
	primes := make([]int, 2)
	primes[0] = 2
	primes[1] = 3
	primes = append(primes, 5)
	primes = append(primes, 7)
	fmt.Println(primes) // => [2 3 5 7]
	// Want to do the same with an array? Have to throw it out and restart with a bigger one.
	//• In most cases you should use slices instead of arrays.
}

func intro() {

	// Array type written as [size]ContainedType
	var myArray [3]string

	// Slice type written as []ContainedType
	var mySlice []string
	mySlice = make([]string, 2)

	// Map type written as map[KeyType]ValueType
	var myMap map[string]int
	myMap = make(map[string]int)


	myArray[0] = "Amy"
	fmt.Println(myArray[0]) // => Amy

	mySlice[1] = "Jose"
	fmt.Println(mySlice[1]) // => Jose

	myMap["Ben"] = 78
	fmt.Println(myMap["Ben"]) // => 78
}

func literals() {

	fmt.Println("\nLiterals and loop\n")
	// Create a collection and add data at the same time.
	myArray := [3]string{"Amy", "Jose", "Ben"}
	mySlice := []string{"Amy", "Jose", "Ben"}
	myMap := map[string]int{"Amy": 84, "Jose": 96, "Ben": 78}
	fmt.Println(myArray[1]) // => Jose
	fmt.Println(mySlice[0]) // => Amy
	fmt.Println(myMap["Ben"]) // => 78

	names := [3]string{"Amy2", "Jose2", "Ben2"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func useRange() {
	fmt.Println("\nRange loop\n")
	nameArray := [3]string{"Amy", "Jose", "Ben"}
	for index, name := range nameArray {
		fmt.Println(index, name)
	}

	// same for slices
	nameSlice := []string{"Amy2", "Jose2", "Ben2"}
	for index, name := range nameSlice {
		fmt.Println(index, name)
	}

	// maps
	grades := map[string]int{"Amy": 84, "Jose": 96, "Ben": 78}
	for name, grade := range grades {
		fmt.Println(name, grade)
	}

	// Don’t want the index, or don’t want the element? Assign it to the blank identifier
	names := [3]string{"Amy3", "Jose3", "Ben3"}
	for _, name := range names {
		fmt.Println(name)
	}
	for index, _ := range names {
		fmt.Println(index)
	}

	for _, name := range names {
		fmt.Println(name)
	}
	for index, _ := range names {
		fmt.Println(index)
	}

	for _, grade := range grades {
		fmt.Println(grade)
	}
	for name, _ := range grades {
		fmt.Println(name)
	}
}

func useStructs() {
	fmt.Println("\nStructs\n")
	// just for one var
	var bucket struct {
		number float64
		word string
		toggle bool
	}
	bucket.number = 3.14
	bucket.word = "pie"
	bucket.toggle = true
	fmt.Println(bucket.number) // => 3.14
	fmt.Println(bucket.word) // => pie
	fmt.Println(bucket.toggle) // => true

	// new data type
	type myType struct {
		number float64
		word string
		toggle bool
	}

	var bucket2 myType
	bucket2.number = 3.1415
	bucket2.word = "pie2"
	bucket2.toggle = false
	fmt.Println(bucket2.number) // => 3.14
	fmt.Println(bucket2.word) // => pie
	fmt.Println(bucket2.toggle) // => true


}

type Coordinates struct {
	Latitude float64
	Longitude float64
}

type Landmark struct {
	Name string
	// An "anonymous field"
	// Has no name of its own, just a type
	Coordinates
}

func useStructs2() {
	var l Landmark
	l.Name = "The Googleplex"
	// Fields for "embedded struct" are "promoted"
	l.Latitude = 37.42
	l.Longitude = -122.08
	fmt.Println(l.Name, l.Latitude, l.Longitude)
}

func main() {

	intro()

	shortDeclarations()
	expandingSlices()
	valuesCanBeOfAnyType()
	literals()
	useRange()
	useStructs()
	useStructs2()

}

/* 
A demo of type assertions.
In this demo we show how to check whether an 
expression of an interface e.g., interF, is
of an interface type e.g., SomeStruct.

*/

package main

import (
	"fmt"
)

// Create interface
var AnInterface interface {
	That()
}

// Create interface compatible types
type SomeStruct struct {
	X int
}

type AnotherStruct struct {
	X int
}

// Create some methods for our new types
func (s SomeStruct) That () {
	fmt.Println("I am ", s)
}

func (s AnotherStruct) That () {
	fmt.Println("I am ", s)
}

func main() {
	// Assign interface
	interF := AnInterface
	// Only add SomeStruct as an interface
	interF = SomeStruct{X: 1}
	// Check if interface expression "interF" is of type SomeStruct
	check1, ok1 := interF.(SomeStruct)
	fmt.Println(check1, ok1)
	// Check if interface expression "interF" is of type AnotherStruct
	check2, ok2 := interF.(AnotherStruct)
	fmt.Println(check2, ok2)

	// Test using if
	if w, ok := interF.(SomeStruct); ok {
		fmt.Println(w, ok)
	}

	// Type switches
	switch x:= interF.(type) {
	case SomeStruct: 
		fmt.Printf("We have a %T\n", x) 
	case AnotherStruct:
		fmt.Printf("We have a %T\n", AnotherStruct{X:1})
	default:
		fmt.Println("We got nothin'\n")
	}
}
// A demo of some type conversions

package main

import (
	"fmt"
	"strconv"
)

var s string
var i int
var f float64

func main() {
	// Print type
	fmt.Printf("Type %T\n", s)
	fmt.Printf("Type %T\n", i)
	fmt.Printf("Type %T\n", f)

	// Convert string to float
	s = "1.2"
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("%g\n", f)

	// Convert string to int
	s = "10"
	i, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("%d\n", i)
}
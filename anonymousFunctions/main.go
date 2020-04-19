// A demo of anonymous functions

package main

import (
	"fmt"
	"strings"
)

// Anonymous functions can only be declared at any level, not just the package level like named functions
// This anon function creates a closure, where the life of the x var persists even after the function is called

// You can declare an anonymous function
var counter = func() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}()

// you can declare anon funcs, in case you need to add recursion
var div func(x, y int) int

func main() {
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2 
	fmt.Println(counter()) // 3

	// You can use anon funcs as args, for example
	x := strings.Map(func(x rune) rune {return x + 1}, "abc" )
	fmt.Println(x)  // bcd 

	// Example of assigning anon function to pre-declared var - useful for recursion
	div = func(x, y int) int {
		if x == 0 {
			x = div(x + 1, y)
		}
		return x + y }
	fmt.Println(div(0, 2))


}


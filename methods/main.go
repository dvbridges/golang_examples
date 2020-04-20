// Demo showing the use of methods 
// See geometry package for method declarations
package main 

import (
	"fmt"
	"methods/geometry"
)

func main() {
	p := geometry.Point{1,2}
	q := geometry.Point{4,6}

	// Use Distance function
	fmt.Println(geometry.Distance(p, q))

	// Use Distance method associated with Point type
	fmt.Println(p.Distance(q))

	// p.Distance is a selector, because it selects the appropriate Distance method
	// for the receiver p of type Point

	// Use another Distance method, but from a different type
	path := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(path.Distance())  // 12
	path = geometry.Path{p, q, p, q}
	fmt.Println(path.Distance())  // 15

	// Using methods with pointer receivers
	// The name of the pointer receiver method is (*Point).ScaleBy
	// The best way to call this is to provide a *Point receiver, like
	r := &geometry.Point{1, 2}
	r.ScaleBy(4)
	fmt.Println(*r)
	// or
	(*r).ScaleBy(4)
	fmt.Println(*r)

}

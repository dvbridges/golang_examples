// A demo: Creating and using structs

package main

import "fmt"

// Structs are aggregate types with n arbitrary types
type Employee struct{
	Name string  // Note, titlecase fields are exportable
	ID int
	Position string
	Salary int
}

// Struct embedding and anonymous fields
type Circle struct {
	X, Y int
}

type Wheel struct {
	Circle  // anon field
	Spokes int
}

func main() {
	jim := Employee{"Jim", 123, "Worker", 10000}
	fmt.Println(jim)

	// Use dot notation to access fields
	fmt.Println(jim.Salary)

	// Give employee a payrise
	jim.Salary += 1000
	fmt.Println(jim.Salary)

	// Get struct pointer
	p := &jim
	// We can write the following to dereference the pointer 
	// i.e., gain access to the value pointer points to 
	fmt.Println((*p).Salary)
	// But that is cumbersome, so Go lets you
	fmt.Println(p.Salary)

	// Give employee a paycut using pointer 
	p.Salary -= 1000
	fmt.Println(jim.Salary)

	// Using struct pointers with functions
	// Useful if structs are large - pass the pointer instead of val
	sackEmployee(&jim)
	fmt.Println(jim.Salary)

	// Example of embedded structs and anonymous fields
	wheel := Wheel{
		Circle: Circle{
			X: 5,
			Y: 5,
		},
		Spokes: 10,
	}
	// Or without explict field assignment
	wheel = Wheel{Circle{1,2},5}

	fmt.Printf("%#v\n", wheel)
	fmt.Printf("%#v\n", wheel.Circle)
	fmt.Printf("%#v\n", wheel.X)
	fmt.Printf("%#v\n", wheel.Y)
	fmt.Printf("%#v\n", wheel.Spokes)

}

func sackEmployee(e *Employee) {
	e.Salary = 0
}
// A demo of variadic functions - packing and unpacking n values into functions

package main

import "fmt"

func main() {

	// you can pass the actual values to a variadic function
	fmt.Println("Expl. sum total: ", sum(1, 2, 3, 4))
	// However, from slices, you must unpack the slice
	// Create slice
	arr := []int{1, 2, 3, 4, 5}
	// Unpack slice into parameter list using ...
	fmt.Println("Slice sum total: ", sum(arr...))
	
	// We can also join strings in the same way
	strarr := []string{"this", "that", "the", "other"}
	fmt.Println(join(strarr...))

}

// sum takes n vals and returns the sum
func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

// Takes n strings and joins them
func join(vals ...string) string {
	str := ""
	for _, v := range vals {
		str += (v + " ")
	}
	return str 
}


// A demo: Using arrays and slices

package main

import (
	"fmt"
)

var arr [5]int
var slices []int

func main() {
	// Ways of creating arrays
	strarr := [...]string{1: "Jan",2: "Feb",3: "Mar", 4: "Apr"}  // Set elements to particular index
	fmt.Println(strarr)
	arr = [5]int{1,2,3,4,5}  // ... for unknown number of elements
	fmt.Println(arr)
	arr := [...]int{1,2,3,4,5}
	fmt.Println(arr)

	// Create empty slice using make and fill
	slices = make([]int, len(arr), cap(arr))
	slices = arr[:]
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

	// Or simply, create slice from array
	slices = arr[:]
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

	// Change length
	slices = slices[:3]
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

	// Change capacity and length - note, changing left number moves slice pointer
	// Once pointer moves, the slice can no longer access array items preceding the pointer
	slices = slices[2:4]
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

	// Show again the capacity and length
	slices = slices[:]
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

	// Change underlying array
	slices[0] = 99
	
	// Loop through array
	for i, v := range arr {
		fmt.Printf("index %d, value %d\n", i, v)
	}

	// Append items using built-in append
	slices = append(slices, 100, 101, 102)
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))
	
	// Append a slice to a slice
	slices = append(slices[:2], slices[:2]...)
	fmt.Println(slices, "Len:", len(slices), "Cap:", cap(slices))

}
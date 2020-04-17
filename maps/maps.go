// A demo: Creating and using maps

package main

import (
	"fmt"
)

func main () {
	// Declare maps using map literals
	first := map[string]int{
	"this": 1,
	}
	for k, v := range first {
		fmt.Printf("key: %s, val: %d\n", k, v)
	}

	// Declare maps using make
	second := make(map[string]int)
	second["that"] = 2
	for k, v := range second {
		fmt.Printf("key: %s, val: %d\n", k, v)
		second[k]++ // increment value for k
		fmt.Printf("key: %s, val: %d\n", k, second[k])
	}
}
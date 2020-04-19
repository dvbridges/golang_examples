// Demo showing the pitfalls of iteration

package main

import "fmt"

var arr [5]int
var storage1 []func()
var storage2 []func()

func main() {
	arr = [5]int{1,2,3,4,5}

	for _, v := range arr {
		// Incorrect method of capturing an iteration variable	
		storage1 = append(storage1, func() { fmt.Println("Oops, incorrect val: ", v * v )})
		// The correct method - assign to new var	
		capturedIter := v
		storage2 = append(storage2, func() { fmt.Println("Correct val: ", capturedIter * capturedIter )})
	}

	for i, _ := range storage1 {
		storage1[i]()  // Has reference to the last iteration var
		storage2[i]()  // Has stored a copy of each iteration var
	}
}

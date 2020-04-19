// A demo looking at panic control

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {

	// Panics occur at runtime
	// We can create a runtime panic by doing things like indexing error,
	// or by creating our own for unexpected situations

	defer printStack()
	createPanic()

	// Note, only use panics when necessary, as they crash the program
	// For diagnostics, dump the stack to the stdout by defering printStack in main

}

func createPanic() {
	smallArr := [5]int{1, 2, 3, 4, 5}
	
	// Create an index error panic
	for i, _ := range smallArr {
		if (i == 2) {
			panic("Index error")
		}
		fmt.Println(smallArr[i + 1])
	}
}

// dump the stack
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
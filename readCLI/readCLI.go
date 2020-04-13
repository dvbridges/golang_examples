// A simple demo showing how to read CLI args

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*

		The simple method is using os package, just like other packages

		fmt.Printf("The current file: %v\n", os.Args[0])
		for i, v := range os.Args[1:] {
			fmt.Printf("Arg %d: %v\n", i, v)
		}

	*/

	// The next method makes use of the bufio package
	// The Scanner handles the split of CLI args and creates an iterator
	input := bufio.NewScanner(os.Stdin)
	count := make(map[string]int)
	exitCount := 0

	fmt.Println(
		`
Your task is to type 5 lines.
Type your lines, and press enter to move on.
If any lines repeat, they will be printed, with their count.
Go`)

	for input.Scan() {
		count[input.Text()]++
		exitCount++

		if exitCount == 5 {
			break
		}
	}

	for k, v := range count {
		if count[k] > 1 {
			fmt.Printf("\n%q repeated %d times\n", k, v)
		}
	}

}

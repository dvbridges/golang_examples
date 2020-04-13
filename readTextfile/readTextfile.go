// Simple demo showing how to read a textfile from CLI and print duplicate
// lines found in the text file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1]           // Get file name from CLI
	count := make(map[string]int) // Create count map

	fileText, err := os.Open(files) // open text file and catch any errors
	if err != nil {
		fmt.Println(err)
	}

	text := bufio.NewScanner(fileText) // Read file text into Scanner
	for text.Scan() {
		count[text.Text()]++
	}

	for k, v := range count { // Print any duplicates
		if count[k] > 1 {
			fmt.Printf("\n%q appears %d times\n", k, v)
		}
	}

	fileText.Close() // close file object

}

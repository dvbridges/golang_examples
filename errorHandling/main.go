// Demo: Showing error handling via propagation
package main

import (
	"fmt"
	"os"
	"time"
	"log"
)

func main() {
	///// Error propagation - propagate errors to caller /////
	s, err := callMe("")  // Raises error if len(s) == 0
	if err != nil {
		fmt.Fprintf(os.Stderr, "callMe ( %q ): %v\n", s, err)
	}
	fmt.Println(s)

	///// Example of a timeout failure /////
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		s, err = callMe("")  // Raises error if len(s) == 0
		if err == nil {
			break
		}
		log.Printf("Call failure: %q; retrying...", err)  // Just log the error
		time.Sleep(time.Second << uint(tries))  // Sleep until next repeat

		if tries >= 2 {
			log.Fatalf("Fatal error: %q;", err)

		}
	}

	fmt.Fprintf(os.Stderr, "Calls to callMe ( %q ): Errors: %v\n", s, err)
	


}

func callMe(s string) (string, error) {
	if len(s) == 0 {
		return s, fmt.Errorf("String has no length: %s", s)
	}

	return s + s, nil
}
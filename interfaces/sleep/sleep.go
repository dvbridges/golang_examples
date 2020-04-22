// Another demo, looking at flags 
package main

import (
	"fmt"
	"flag"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...\n", *period)
	time.Sleep(*period)
}
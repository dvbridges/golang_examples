package main

import "encaps/encaps"

func main() {
	p := encaps.Point{1,2}
	// Distance is exported
	p.Distance(encaps.Point{1,5})
	// printVal is not, and thus private
	p.printVal(4) // throws error
}
// Demonstration of encapsulation

package encaps 

import (
	"fmt"
	"math"
)

// Go uses capitalized identifiers to indicate something is exportable / public
// and non-capitalized idenfifiers to indicate something is not exportable / private

type Point struct {
	X, Y float64 
}

// A public method
func (p Point) Distance(q Point) float64 {
	val := math.Hypot(q.Y-p.Y, q.X-p.X)
	p.printVal(val)
	return val
}

// A private method
func (p Point) printVal (v float64) {
	fmt.Println("The distance is", v)
}
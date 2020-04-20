// Demo showing the use of methods
/*
Looks at:
	Functions
	Methods and naming
	Methods with pointer receivers
*/
package geometry

import "math"

// The method format is:
// 		func (receiver of method) name (parameter-list) results { method body}

// In Go, a method is a function associated with a particular type
type Point struct { X, Y float64 }

// All methods of the same type must have different names
// However, different methods can share the same name, as long as they
// are assigned to different types

// Lets create another type, to assign a new Distance method
type Path []Point

// Traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance as a Point method
func (p Point) Distance (q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance returns the distance travelled along n paths
func (path Path) Distance () float64 {
	sum := 0.0
	for i, _ := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum 
}

// Methods with pointer receivers
// Note, convention dictates that if a type uses a pointer receiver in a method,
// then all methods of that type should use pointer receivers.
func (p *Point) ScaleBy (factor float64) {
	p.X *= factor
	p.Y *= factor
}
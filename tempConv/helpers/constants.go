package tempConv

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

// A Celcius method - returns formatted val as string
func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

// A Fahrenheit method - returns formatted val as string
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}
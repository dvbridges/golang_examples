// Some helper functions

package tempConv

// Convert Celsius to Fahrenheit
func CtoF(c float64) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

// Convert Fahrenheit to Celsius
func FtoC(f float64) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
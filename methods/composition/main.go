// Demonstration of composing types using Struct embedding
package main

import (
	"fmt"
	"composer/composer"
	"image/color"
)

func main() {
	/* ColoredPoint is composed of:
		1) A Point type, a struct of X, Y fields
		2) a color.RGBA type
		3) All methods associated with embedded Point type
	*/
	col := composer.ColoredPoint{composer.Point{1,2}, color.RGBA{255, 0, 0, 255}}
	fmt.Println(col.X, col.Y)
	fmt.Println(col.Color)
	fmt.Println(col.Distance(composer.Point{1,5}))

}
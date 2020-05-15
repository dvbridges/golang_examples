// Bar chart demo

package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"

	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func main() {
	// Create plot
	plt, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create random data values
	groupA := make(plotter.Values, 5)
	groupB := make(plotter.Values, 5)
	var i int
	for i < 5 {
		rand.Seed(int64(i))
		groupA[i] = math.Abs(rand.NormFloat64())
		groupB[i] = math.Abs(rand.NormFloat64())
		i++
	}

	// Create bar
	barWidth := vg.Length(15)
	barA, err := plotter.NewBarChart(groupA, barWidth)
	barB, err := plotter.NewBarChart(groupB, barWidth)

	// Format bar
	barA.Color = color.RGBA{R: 100, G: 150, B: 0, A: 100}
	barB.Color = color.RGBA{R: 150, G: 100, B: 0, A: 100}
	barA.Offset = barWidth

	plt.Add(barA, barB)
	plt.Legend.Add("1", barA)
	plt.Legend.Add("2", barB)
	plt.Legend.Top = true
	plt.NominalX("A", "B", "C", "D", "E")
	plt.Y.Label.Text = "Amount"
	plt.X.Label.Text = "Condition"

	// Write the plot to file
	f, err := os.Create("BarChart.png")
	defer f.Close()
	// Create plot writer
	wt, err := plt.WriterTo(256, 256, "png")
	wt.WriteTo(f)
}

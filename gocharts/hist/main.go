// Histogram demo

package main

import (
	"image/color"
	"log"
	"math/rand"

	"gonum.org/v1/gonum/stat/distuv"
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
	groupA := make(plotter.Values, 1000)
	var i int
	for i < 1000 {
		rand.Seed(int64(i))
		groupA[i] = rand.NormFloat64()
		i++
	}

	// Create bar
	bins := 15
	hist, err := plotter.NewHist(groupA, bins)
	hist.Normalize(1)

	// Format bar
	hist.FillColor = color.RGBA{R: 100, G: 150, B: 0, A: 100}
	hist.LineStyle.Width = vg.Length(.45)

	// Fit the normal distribution function
	norm := plotter.NewFunction(distuv.UnitNormal.Prob)
	norm.Color = color.RGBA{G: 255, A: 255}
	norm.Width = vg.Points(2)
	norm.Dashes = []vg.Length{vg.Points(10), vg.Points(1)}

	plt.Add(hist, norm)
	plt.Title.Text = "Histogram"
	plt.Y.Label.Text = "Density"
	plt.X.Label.Text = "Bins"
	plt.Y.Max = .4
	plt.X.Min, plt.X.Max = -4, 4

	// Write the plot to file
	// f, err := os.Create("Hist.png")
	// defer f.Close()
	// Create plot writer
	err = plt.Save(256, 256, "Histo.png")
	if err != nil {
		panic(err)
	}
	// wt, err := plt.WriterTo(256, 256, "png")
	// wt.WriteTo(f)
}

// A real time chart demo using gonum

package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

type server struct {
	data []time.Duration
	sync.RWMutex
}

func main() {
	rand.Seed(time.Now().Unix())

	var s server

	http.HandleFunc("/", s.root)
	http.HandleFunc("/statz", s.statz)
	http.HandleFunc("/statz/scatter.png", s.scatter)
	http.HandleFunc("/statz/hist.png", s.hist)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func (s *server) root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	x := 1000 * rand.NormFloat64()
	d := time.Duration(x) * time.Millisecond
	time.Sleep(10)
	fmt.Fprintln(w, "sleep for", d)

	s.Lock()
	s.data = append(s.data, d)

	if len(s.data) > 1000 {
		s.data = s.data[len(s.data)-1000:]
	}
	s.Unlock()
}

func (s *server) statz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s",
		`<h1>Latency Stats</h1>
		<div style="display:flex">
		<img src="/statz/scatter.png?rand=0" style="width:30%">
		<img src="/statz/hist.png?rand=0" style="width:30%">
		</div>
		<script>
			setInterval(function() {
				var imgs = document.getElementsByTagName("img")
				for (let i=0; i<imgs.length; i++) {
					src = imgs[i].src.split("?")[0]
					newRand = Math.random()
					imgs[i].src = src + "?" + newRand
				}
			}, 250)
		</script>
		`)
}

func (s *server) hist(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	// Create plot for scatter
	plt, err := plot.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vs := make(plotter.Values, len(s.data))
	for i, v := range s.data {
		vs[i] = float64(v) / float64(time.Millisecond)
	}
	h, err := plotter.NewHist(vs, 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// format plot
	plt.Add(h)
	plt.Title.Text = "Latency distribution"
	plt.Y.Label.Text = "Milliseconds"

	// Create writer
	wt, err := plt.WriterTo(256, 256, "png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "image/png")
	wt.WriteTo(w)

}

func (s *server) scatter(w http.ResponseWriter, r *http.Request) {
	s.RLock()
	defer s.RUnlock()

	// Make a slice containing len(data) XYs types
	xys := make(plotter.XYs, len(s.data))
	for i, d := range s.data {
		xys[i].X = float64(i)
		xys[i].Y = float64(d) / float64(time.Millisecond)
	}

	// Create scatter
	scat, err := plotter.NewScatter(xys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create mean line
	avgs := make(plotter.XYs, len(s.data))
	var sum float64
	for i, d := range s.data {
		sum += float64(d)
		avgs[i].X = float64(i)
		avgs[i].Y = sum / (float64(time.Millisecond) * float64(i+1))
	}
	lin, err := plotter.NewLine(avgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create basic grid
	grd := plotter.NewGrid()

	// Create plot for scatter
	plt, err := plot.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// format plot
	plt.Add(scat, lin, grd)
	plt.Title.Text = "Endpoint Latency"
	plt.Y.Label.Text = "Milliseconds"
	plt.X.Label.Text = "Sample"
	// Format data points
	scat.GlyphStyle.Shape = draw.CrossGlyph{}
	lin.Color = color.RGBA{G: 150, A: 150}
	grd.Horizontal.Color = color.RGBA{R: 150, A: 150}
	grd.Vertical.Width = 0

	// Create writer
	wt, err := plt.WriterTo(256, 256, "png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("Content-Type", "image/png")
	wt.WriteTo(w)

}

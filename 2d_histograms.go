package main

import (
	"log"
	"math"

	"gocv.io/x/gocv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// HistGrid is a custom type that stores the 2D histogram and implements plotter.GridXYZ.
type HistGrid struct {
	hist *gocv.Mat
	rows int
	cols int
}

// Dims returns the dimensions of the grid (rows and columns).
func (hg HistGrid) Dims() (int, int) {
	return hg.rows, hg.cols
}

// Z returns the value at a given (i, j) position in the grid.
func (hg HistGrid) Z(i, j int) float64 {
	return float64(hg.hist.GetFloatAt(i, j))
}

// X returns the X coordinate for the given column index (j).
func (hg HistGrid) X(j int) float64 {
	return float64(j) * 8 // Scale back to the original intensity range (0-255)
}

// Y returns the Y coordinate for the given row index (i).
func (hg HistGrid) Y(i int) float64 {
	return float64(i) * 8 // Scale back to the original intensity range (0-255)
}

func main() {
	// Load the color image
	img := gocv.IMRead("images/city_hall.jpg", gocv.IMReadColor)
	if img.Empty() {
		log.Fatalf("failed to read image")
	}
	defer img.Close()

	// Split the image into color channels (B, G, R)
	chans := gocv.Split(img)
	blueChannel := chans[0]  // Blue channel
	greenChannel := chans[1] // Green channel

	// Calculate the 2D histogram for Green and Blue channels
	hist := gocv.NewMat()
	defer hist.Close()
	gocv.CalcHist([]gocv.Mat{greenChannel, blueChannel}, []int{0, 1}, gocv.NewMat(), &hist, []int{32, 32}, []float64{0, 256, 0, 256}, false)

	// Create a plot and set title and labels
	p := plot.New()
	p.Title.Text = "2D Green-Blue Histogram"
	p.X.Label.Text = "Green Intensity"
	p.Y.Label.Text = "Blue Intensity"

	// Create a HistGrid struct that holds the histogram and grid dimensions
	histGrid := HistGrid{hist: &hist, rows: 32, cols: 32}

	// Normalize the histogram values using logarithmic scale
	maxValue := maxHistogramValue(histGrid)
	for i := 0; i < histGrid.rows; i++ {
		for j := 0; j < histGrid.cols; j++ {
			val := histGrid.Z(i, j)
			hist.SetFloatAt(i, j, float32(math.Log(1+val)/math.Log(1+maxValue)))
		}
	}

	// Create a heatmap and add it to the plot
	heatMap := plotter.NewHeatMap(histGrid, palette.Heat(10, 1))
	p.Add(heatMap)

	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "images/2d_green_blue_histogram.png"); err != nil {
		log.Fatalf("could not save plot: %v", err)
	}
	h := gocv.IMRead("images/2d_green_blue_histogram.png", gocv.IMReadColor)
	defer h.Close()
	win := gocv.NewWindow("histogram")
	defer win.Close()
	win.IMShow(h)
	gocv.WaitKey(0)
}

// maxHistogramValue finds the maximum Z value in the histogram to normalize the heatmap values.
func maxHistogramValue(hg HistGrid) float64 {
	max := 0.0
	for i := 0; i < hg.rows; i++ {
		for j := 0; j < hg.cols; j++ {
			val := hg.Z(i, j)
			if val > max {
				max = val
			}
		}
	}
	return max
}

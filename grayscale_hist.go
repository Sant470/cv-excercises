package main

import (
	"image/color"
	"log"

	"gocv.io/x/gocv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)
	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
	defer img.Close()

	// calculate hist
	hist := gocv.NewMat()
	gocv.CalcHist([]gocv.Mat{img}, []int{0}, gocv.NewMat(), &hist, []int{256}, []float64{0, 256}, false)

	// Extract the histogram values
	histValues := make(plotter.Values, 256)
	for i := 0; i < 256; i++ {
		histValues[i] = float64(hist.GetFloatAt(i, 0))
	}

	// Plot the histogram using gonum/plot
	p := plot.New()
	p.Title.Text = "Grayscale Histogram"
	p.X.Label.Text = "Pixel Intensity"
	p.Y.Label.Text = "Frequency"

	// Create a bar plot
	bars, err := plotter.NewBarChart(histValues, vg.Points(2))
	if err != nil {
		log.Fatalf("failed to create bar chart: %v", err)
	}

	bars.LineStyle.Width = vg.Length(0)
	bars.Color = color.Gray{Y: 128}

	// Add bars to the plot
	p.Add(bars)
	// Save the plot to a PNG file
	if err := p.Save(6*vg.Inch, 4*vg.Inch, "images/grayscale_histogram.png"); err != nil {
		log.Fatalf("failed to save plot: %v", err)
	}

	h := gocv.IMRead("images/grayscale_histogram.png", gocv.IMReadGrayScale)
	defer h.Close()
	win := gocv.NewWindow("histogram")
	defer win.Close()
	win.IMShow(h)
	gocv.WaitKey(0)
}

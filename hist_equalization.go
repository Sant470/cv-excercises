/*
Histogram equalization imporoves the contrast of an image by stretching the distribution of pixels.
*/

package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadGrayScale)
	defer img.Close()

	// create histogram equalization
	eq := gocv.NewMat()
	gocv.EqualizeHist(img, &eq)

	// Save the equalized image to a file
	outputPath := "images/equalized_hist.jpg"
	if ok := gocv.IMWrite(outputPath, eq); !ok {
		fmt.Println("Error saving image")
	}

	// display the images ...
	hstack := gocv.NewMat()
	defer hstack.Close()
	gocv.Hconcat(img, eq, &hstack)
	win := gocv.NewWindow("hist equalization")
	win.IMShow(hstack)
	gocv.WaitKey(0)
}

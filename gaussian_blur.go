/*
Practically, this means that each pixel in the image is mixed in with its surrounding pixel intensities.
This "mixture" of pixels in a neighborhood becomes our blurred pixel.
*/

package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {

	// Gaussian blur is similar to average blurring, but instead of using simple mean,
	// we are now using weighted mean, where neighborhood pixels that are closer to to the central pixel
	// contribute more weight to the average.

	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)

	hstack := gocv.NewMat()
	blur3 := gocv.NewMat()
	blur5 := gocv.NewMat()
	blur7 := gocv.NewMat()

	defer hstack.Close()
	defer blur3.Close()
	defer blur5.Close()
	defer blur7.Close()

	gocv.GaussianBlur(img, &blur3, image.Point{3, 3}, 0, 0, gocv.BorderConstant)
	gocv.GaussianBlur(img, &blur5, image.Point{5, 5}, 0, 0, gocv.BorderConstant)
	gocv.GaussianBlur(img, &blur7, image.Point{7, 7}, 0, 0, gocv.BorderConstant)
	gocv.Hconcat(blur3, blur5, &hstack)
	gocv.Hconcat(hstack, blur7, &hstack)

	win1 := gocv.NewWindow("gaussian blur")
	win1.IMShow(hstack)
	gocv.WaitKey(0)
}

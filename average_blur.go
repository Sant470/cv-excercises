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

	// Averaging
	// We are going to define a k*k sliding window on top of our image, where k is always an odd number.
	// This window slide from left to right and from top to bottom. The pixel at the centre of this matrix
	// is set to the average of all other pixels surrounding it.

	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)

	hstack := gocv.NewMat()
	blur3 := gocv.NewMat()
	blur5 := gocv.NewMat()
	blur7 := gocv.NewMat()

	defer hstack.Close()
	defer blur3.Close()
	defer blur5.Close()
	defer blur7.Close()

	gocv.Blur(img, &blur3, image.Point{3, 3})
	gocv.Blur(img, &blur5, image.Point{5, 5})
	gocv.Blur(img, &blur7, image.Point{7, 7})
	gocv.Hconcat(blur3, blur5, &hstack)
	gocv.Hconcat(hstack, blur7, &hstack)

	win1 := gocv.NewWindow("average blur")
	win1.IMShow(hstack)
	gocv.WaitKey(0)
}

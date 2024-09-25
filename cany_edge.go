/*
The canny edge detector is multi-step process. It involves blurring the image to remove noise, computing sobel gradient image in
the x and y direction, supressing edges, and finally a hysteresis thresholding stage that determine if a pixel is "edge-like" or not.
*/

package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/coins.jpeg", gocv.IMReadGrayScale)
	blur := gocv.NewMat()
	defer img.Close()
	defer blur.Close()
	gocv.GaussianBlur(img, &blur, image.Pt(5, 5), 0, 0, gocv.BorderDefault)

	win1 := gocv.NewWindow("blur image")
	win1.IMShow(blur)
	win1.WaitKey(0)

	canny := gocv.NewMat()
	defer canny.Close()
	// Gradient values between 30 and 155 are either classified as edges or non-edges based on how their intensities are "connected".
	gocv.Canny(blur, &canny, 30, 155)
	win2 := gocv.NewWindow("canny")
	win2.IMShow(canny)
	win2.WaitKey(0)
}

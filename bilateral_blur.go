/*
Bilateral filtering considers both the spatial proximity of pixels and the intensity differences between them. It works by combining two Gaussian filters:
Spatial Gaussian: This filter ensures that only nearby pixels contribute to the blur, based on their spatial distance.
Intensity (Range) Gaussian: This filter ensures that only pixels with similar intensity (color or brightness) contribute to the blur.
In essence, bilateral filtering assigns more weight to pixels that are both spatially close and have similar intensity, leading to a selective blurring effect that preserves sharp edges
*/

package main

import (
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

	gocv.BilateralFilter(img, &blur3, 5, 21, 21)
	gocv.BilateralFilter(img, &blur5, 7, 31, 31)
	gocv.BilateralFilter(img, &blur7, 9, 41, 41)
	gocv.Hconcat(blur3, blur5, &hstack)
	gocv.Hconcat(hstack, blur7, &hstack)

	win1 := gocv.NewWindow("bilateral filter")
	win1.IMShow(hstack)
	gocv.WaitKey(0)
}

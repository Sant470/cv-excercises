/*
Adaptive thresholding considers small neighbors of pixels and then finds an optimal value T for each neighbor. This method allow us
to handle cases where there may be dramatic ranghes of pixel intensities and the optimal value of T may change for different parts of the image.
*/

package main

import (
	"image"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/coins.jpeg", gocv.IMReadGrayScale)
	blurImg := gocv.NewMat()
	gocv.GaussianBlur(img, &blurImg, image.Point{5, 5}, 0, 0, gocv.BorderConstant)

	defer img.Close()
	defer blurImg.Close()

	win1 := gocv.NewWindow("grayscale image")
	win1.IMShow(blurImg)
	win1.WaitKey(0)

	thresholdImg1 := gocv.NewMat()
	defer thresholdImg1.Close()
	gocv.AdaptiveThreshold(blurImg, &thresholdImg1, 255, gocv.AdaptiveThresholdMean, gocv.ThresholdBinaryInv, 11, 4)
	win2 := gocv.NewWindow("threshold adaptive mean")
	win2.IMShow(thresholdImg1)
	win2.WaitKey(0)

	thresholdImg2 := gocv.NewMat()
	defer thresholdImg2.Close()
	gocv.AdaptiveThreshold(blurImg, &thresholdImg2, 255, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinaryInv, 15, 3)
	win3 := gocv.NewWindow("Gaussian")
	win3.IMShow(thresholdImg2)
	win3.WaitKey(0)

}

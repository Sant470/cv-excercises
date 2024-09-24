/*
Thresholding is the binarization of an image. In general, we seek to convert a grayscale image to a binary image,
where the pixels are either 0 or 255.

A simple thresholding example would be selecting a pixel value p, and then setting all pixel intensities less than p to zero, and all
pixel value greater than p to 255.
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

	thresholdImg := gocv.NewMat()
	defer thresholdImg.Close()
	gocv.Threshold(blurImg, &thresholdImg, 155, 255, gocv.ThresholdOtsu)
	win2 := gocv.NewWindow("thresholdImg")
	win2.IMShow(thresholdImg)
	win2.WaitKey(0)

	thresholdInv := gocv.NewMat()
	defer thresholdInv.Close()
	gocv.Threshold(blurImg, &thresholdInv, 155, 255, gocv.ThresholdBinaryInv)
	win3 := gocv.NewWindow("threshold inv")
	win3.IMShow(thresholdInv)
	win3.WaitKey(0)

	bitwiseAND := gocv.NewMat()
	defer bitwiseAND.Close()
	gocv.BitwiseAnd(img, thresholdInv, &bitwiseAND)
	win4 := gocv.NewWindow("bitwise and")
	win4.IMShow(bitwiseAND)
	win4.WaitKey(0)

}

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

	gocv.GaussianBlur(img, &blur, image.Point{5, 5}, 0, 0, gocv.BorderDefault)
	win1 := gocv.NewWindow("gray scale")
	win1.IMShow(blur)
	win1.WaitKey(0)
	defer win1.Close()

	// constants ...
	kernelSize := 3
	scale := 1.0
	delta := 0.0

	// laplacian gradient
	lap := gocv.NewMat()
	defer lap.Close()
	gocv.Laplacian(blur, &lap, gocv.MatTypeCV64F, kernelSize, scale, delta, gocv.BorderDefault)
	win2 := gocv.NewWindow("laplacian gradient")
	win2.IMShow(lap)
	win2.WaitKey(0)
	defer win2.Close()
	// sobel gradient
	sobelX := gocv.NewMat()
	sobelY := gocv.NewMat()
	sobel := gocv.NewMat()

	defer sobelX.Close()
	defer sobelY.Close()
	defer sobel.Close()

	gocv.Sobel(blur, &sobelX, gocv.MatTypeCV64F, 1, 0, kernelSize, scale, delta, gocv.BorderDefault)
	gocv.Sobel(blur, &sobelY, gocv.MatTypeCV64F, 0, 1, kernelSize, scale, delta, gocv.BorderDefault)
	gocv.BitwiseOr(sobelX, sobelY, &sobel)

	win3 := gocv.NewWindow("sobel x")
	win4 := gocv.NewWindow("sobel y")
	win5 := gocv.NewWindow("sobel")

	defer win3.Close()
	defer win4.Close()
	defer win5.Close()

	win3.IMShow(sobelX)
	win3.WaitKey(0)

	win4.IMShow(sobelY)
	win4.WaitKey(0)

	win5.IMShow(sobel)
	win5.WaitKey(0)
}

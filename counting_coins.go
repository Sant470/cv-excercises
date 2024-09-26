/*
A contour is a curve of points, with no gaps in the curve.
*/

package main

import (
	"fmt"
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/coins.jpeg", gocv.IMReadGrayScale)
	defer img.Close()

	// gaussian blur
	blur := gocv.NewMat()
	defer blur.Close()
	gocv.GaussianBlur(img, &blur, image.Pt(5, 5), 0, 0, gocv.BorderDefault)
	win := gocv.NewWindow("img")
	win.IMShow(img)
	win.WaitKey(0)

	// edge
	edge := gocv.NewMat()
	defer edge.Close()
	gocv.Canny(img, &edge, 30, 155)
	win1 := gocv.NewWindow("edges")
	win1.IMShow(edge)
	win1.WaitKey(0)

	// contours
	cnts := gocv.FindContours(edge.Clone(), gocv.RetrievalExternal, gocv.ChainApproxSimple)
	defer cnts.Close()
	fmt.Printf("contours %#v\n", cnts.Size())

	coins := img.Clone()
	gocv.DrawContours(&coins, cnts, -1, color.RGBA{0, 255, 0, 0}, 1)
	win2 := gocv.NewWindow("contours")
	win2.IMShow(coins)
	win2.WaitKey(0)

	// croping each individual coin from the image
	for i := 0; i < cnts.Size(); i++ {
		rect := gocv.BoundingRect(cnts.At(i))
		coin := img.Region(rect)
		coinWindow := gocv.NewWindow(fmt.Sprintf("Coin %d", i+1))
		defer coinWindow.Close()

		coinWindow.IMShow(coin)
		gocv.IMWrite(fmt.Sprintf("images/coin_%d.jpeg", i+1), coin) // Save each cropped coin as an image
		coinWindow.WaitKey(0)
		// Close the cropped coin image
		coin.Close()
	}

}

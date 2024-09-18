/*
A mask allow us to focus only on the portion of the image that interest us.
*/

package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)
	defer img.Close()
	win1 := gocv.NewWindow("original image")
	win1.IMShow(img)
	win1.WaitKey(0)

	// creating a mask
	rows := img.Size()[0]
	cols := img.Size()[1]
	cx := cols / 2
	cy := rows / 2
	minPoint := image.Point{cx - 200, cy - 400}
	maxPoint := image.Point{cx + 200, cy}
	rect := image.Rectangle{minPoint, maxPoint}
	mask := gocv.NewMatWithSize(rows, cols, gocv.MatTypeCV8U)
	defer mask.Close()

	gocv.Rectangle(&mask, rect, color.RGBA{255, 255, 255, 0}, -1)
	mask3C := gocv.NewMat()
	defer mask3C.Close()
	gocv.Merge([]gocv.Mat{mask, mask, mask}, &mask3C)
	win2 := gocv.NewWindow("mask")
	win2.IMShow(mask)
	win2.WaitKey(0)

	// creating masked image
	masked := gocv.NewMat()
	defer masked.Close()
	gocv.BitwiseAnd(img, mask3C, &masked)
	win3 := gocv.NewWindow("masked")
	win3.IMShow(masked)
	win3.WaitKey(0)
}

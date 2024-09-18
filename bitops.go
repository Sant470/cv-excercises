package main

import (
	"image"
	"image/color"

	"gocv.io/x/gocv"
)

func main() {

	// draw rectangle ...
	img := gocv.NewMatWithSize(400, 400, gocv.MatTypeCV8U)
	minPoint := image.Point{25, 25}
	maxPoint := image.Point{275, 275}
	rect := image.Rectangle{minPoint, maxPoint}
	gocv.Rectangle(&img, rect, color.RGBA{255, 255, 255, 1}, -1)
	win2 := gocv.NewWindow("rect")
	win2.IMShow(img)
	win2.WaitKey(0)

	// draw circle
	img2 := gocv.NewMatWithSize(400, 400, gocv.MatTypeCV8U)
	gocv.Circle(&img2, image.Point{150, 150}, 150, color.RGBA{255, 255, 255, 1}, -1)
	win3 := gocv.NewWindow("circle")
	win3.IMShow(img2)
	win3.WaitKey(0)

	// bitwise AND
	bitwiseAND := gocv.NewMat()
	gocv.BitwiseAnd(img, img2, &bitwiseAND)
	win4 := gocv.NewWindow("bitwiseAND")
	win4.IMShow(bitwiseAND)
	win4.WaitKey(0)

	// bitwise OR
	bitwiseOR := gocv.NewMat()
	gocv.BitwiseOr(img, img2, &bitwiseOR)
	win5 := gocv.NewWindow("bitwiseOR")
	win5.IMShow(bitwiseOR)
	win5.WaitKey(0)

	// bitwise XOR
	bitwiseXOR := gocv.NewMat()
	gocv.BitwiseXor(img, img2, &bitwiseXOR)
	win6 := gocv.NewWindow("bitwiseXOR")
	win6.IMShow(bitwiseXOR)
	win6.WaitKey(0)

	// bitwise NOT
	bitwiseNOT := gocv.NewMat()
	gocv.BitwiseNot(img, &bitwiseNOT)
	win7 := gocv.NewWindow("bitwise NOT")
	win7.IMShow(bitwiseNOT)
	win7.WaitKey(0)

}

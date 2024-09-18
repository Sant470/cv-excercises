package main

import (
	"github.com/sant470/cv-basic/utils"
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)
	win := gocv.NewWindow("original")
	win.IMShow(img)
	win.WaitKey(0)

	// translate left 100, up 100
	translatedImg := utils.Resize(&img, 100, 0, gocv.InterpolationArea)
	win1 := gocv.NewWindow("translated image")
	win1.IMShow(translatedImg)
	win1.WaitKey(0)

	// translate right 100 and down 100
	translatedImg1 := utils.Resize(&img, 40, 0, gocv.InterpolationArea)
	win2 := gocv.NewWindow("translated image 2")
	win2.IMShow(translatedImg1)
	win2.WaitKey(0)
}

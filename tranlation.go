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
	translatedImg := utils.Translate(&img, -100, -100)
	win1 := gocv.NewWindow("translated image")
	win1.IMShow(translatedImg)
	win1.WaitKey(0)

	// translate right 100 and down 100
	translatedImg1 := utils.Translate(&img, 100, 100)
	win2 := gocv.NewWindow("translated image 2")
	win2.IMShow(translatedImg1)
	win2.WaitKey(0)
}

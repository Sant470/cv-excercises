package main

import (
	"github.com/sant470/cv-basic/utils"
	"gocv.io/x/gocv"
)

const (
	_scale = 1
)

func main() {
	img := gocv.IMRead("images/city_hall.jpg", gocv.IMReadColor)
	win := gocv.NewWindow("original")
	win.IMShow(img)
	win.WaitKey(0)

	// translate left 100, up 100
	translatedImg := utils.Rotate(&img, nil, 30, 1)
	win1 := gocv.NewWindow("translated image")
	win1.IMShow(translatedImg)
	win1.WaitKey(0)

	// translate right 100 and down 100
	translatedImg1 := utils.Rotate(&img, nil, -30, 1)
	win2 := gocv.NewWindow("translated image 2")
	win2.IMShow(translatedImg1)
	win2.WaitKey(0)

	// translate right 100 and down 100
	translatedImg2 := utils.Rotate(&img, nil, -180, 1)
	win3 := gocv.NewWindow("translated image 3")
	win3.IMShow(translatedImg2)
	win3.WaitKey(0)

}

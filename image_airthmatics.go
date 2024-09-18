package main

import "gocv.io/x/gocv"

func main() {
	img := gocv.IMRead("images/sunflower.jpg", gocv.IMReadColor)
	h, w := img.Rows(), img.Cols()
	win := gocv.NewWindow("original")
	win.IMShow(img)
	win.WaitKey(0)

	// adding 100 pixels globally
	m1 := gocv.NewMatWithSize(h, w, gocv.MatTypeCV8UC3)
	m1.SetTo(gocv.NewScalar(100, 100, 100, 0)) // 100 for all channels except alpha
	dst1 := gocv.NewMat()
	gocv.Add(img, m1, &dst1)
	win2 := gocv.NewWindow("brighter image")
	win2.IMShow(dst1)
	win2.WaitKey(0)

	// subtract 50 pixels globally
	m2 := gocv.Ones(h, w, gocv.MatTypeCV8UC3)
	m2.SetTo(gocv.NewScalar(50, 50, 50, 0))
	dst2 := gocv.NewMat()
	gocv.Subtract(img, m2, &dst2)
	win3 := gocv.NewWindow("darker image")
	win3.IMShow(dst2)
	win3.WaitKey(0)
}

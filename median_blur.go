/*
When applying a median blur, we define our kernel size k. Then, as in the averaging blurring method,
we consider all pixels in the neighborhood of size k*k. But, unlike the averaging method, instead of
replacing the central pixel with the average of the neighborhood, we instead replace the central pixel with
the cental pixel with the median of tghe neighborhood.
*/

package main

import (
	"gocv.io/x/gocv"
)

func main() {

	img := gocv.IMRead("images/color-salt-and-pepper-noise.png", gocv.IMReadColor)

	hstack := gocv.NewMat()
	blur3 := gocv.NewMat()
	blur5 := gocv.NewMat()
	blur7 := gocv.NewMat()

	defer hstack.Close()
	defer blur3.Close()
	defer blur5.Close()
	defer blur7.Close()

	gocv.MedianBlur(img, &blur3, 3)
	gocv.MedianBlur(img, &blur5, 5)
	gocv.MedianBlur(img, &blur7, 7)

	gocv.Hconcat(blur3, blur5, &hstack)
	gocv.Hconcat(hstack, blur7, &hstack)

	win1 := gocv.NewWindow("median blur")
	win1.IMShow(hstack)
	gocv.WaitKey(0)
}

package utils

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func Translate(img *gocv.Mat, tx, ty float64) gocv.Mat {
	m := gocv.NewMatWithSize(2, 3, gocv.MatTypeCV64F)
	m.SetDoubleAt(0, 0, 1.0)
	m.SetDoubleAt(0, 1, 0.0)
	m.SetDoubleAt(0, 2, tx)
	m.SetDoubleAt(1, 0, 0.0)
	m.SetDoubleAt(1, 1, 1.0)
	m.SetDoubleAt(1, 2, ty)
	defer m.Close()
	dim := img.Size()
	pt := image.Point{dim[1], dim[0]}
	result := gocv.NewMat()
	gocv.WarpAffine(*img, &result, m, pt)
	return result
}

func Rotate(img *gocv.Mat, centre []int, angle float64, scale float64) gocv.Mat {
	dim := img.Size()
	if centre == nil {
		x := dim[1] / 2
		y := dim[0] / 2
		centre = append(centre, x, y)
	}
	cp := image.Point{centre[0], centre[1]}
	rotationMatrix := gocv.GetRotationMatrix2D(cp, angle, 1)
	defer rotationMatrix.Close()
	result := gocv.NewMat()
	gocv.WarpAffine(*img, &result, rotationMatrix, image.Point{dim[1], dim[0]})
	return result
}

func Resize(img *gocv.Mat, height, width int, flag gocv.InterpolationFlags) gocv.Mat {
	if height == 0 && width == 0 {
		return *img
	}
	dim := img.Size()
	h, w := dim[0], dim[1]
	if width == 0 {
		r := float64(height) / float64(h)
		width = int(r * float64(w))
	}
	if height == 0 {
		r := float64(width) / float64(w)
		height = int(r * float64(h))
	}
	fmt.Println("height: ", height, "width: ", width)
	result := gocv.NewMat()
	gocv.Resize(*img, &result, image.Point{width, height}, 0, 0, flag)
	return result
}

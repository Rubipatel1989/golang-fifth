package utils

import (
	"fmt"

	"gocv.io/x/gocv"
)

func CompareFaces(profilePath, idPath string) float64 {
	img1 := gocv.IMRead(profilePath, gocv.IMReadGrayScale)
	defer img1.Close()

	img2 := gocv.IMRead(idPath, gocv.IMReadGrayScale)
	defer img2.Close()

	if img1.Empty() || img2.Empty() {
		fmt.Println("Error loading images")
		return 0
	}

	result := gocv.NewMat()
	defer result.Close()

	gocv.MatchTemplate(img1, img2, &result, gocv.TmCcoeffNormed, gocv.NewMat())

	var maxVal float32
	_, maxVal, _, _ = gocv.MinMaxLoc(result)

	return float64(maxVal)
}

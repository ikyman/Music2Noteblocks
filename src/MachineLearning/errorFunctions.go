package mCraftnBlockmLearning

import (
	"math"
	"gonum.org/v1/gonum/mat"
)

func rmseError(estimated mat.Matrix, known mat.Matrix) float64 {
	er, ec := estimated.Dims()
	kr, kc := known.Dims()
	if er != kr || ec != kc {
		panic("estimated and known must have the same dimensions")
	}

	sumSq := 0.0
	for r := 0; r < er; r++ {
		for c := 0; c < ec; c++ {
			diff := estimated.At(r, c) - known.At(r, c)
			sumSq += diff * diff
		}
	}

	n := float64(er * ec)
	if n == 0 {
		return 0
	}
	return math.Sqrt(sumSq / n)
}

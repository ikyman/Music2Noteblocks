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

// Element-wise d(RMSE)/d(estimated). Same shape as estimated/known.
func rmseErrorDerivative(estimated mat.Matrix, known mat.Matrix) mat.Matrix {
	er, ec := estimated.Dims()
	kr, kc := known.Dims()
	if er != kr || ec != kc {
		panic("estimated and known must have the same dimensions")
	}

	n := float64(er * ec)
	out := mat.NewDense(er, ec, nil)
	if n == 0 {
		return out
	}

	rmse := rmseError(estimated, known)
	if rmse == 0 {
		return out
	}

	scale := 1.0 / (n * rmse)
	for r := 0; r < er; r++ {
		for c := 0; c < ec; c++ {
			out.Set(r, c, (estimated.At(r, c)-known.At(r, c))*scale)
		}
	}
	return out
}

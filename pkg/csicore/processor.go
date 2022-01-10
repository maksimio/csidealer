package csicore

import "math/cmplx"

func CsiToAbs(csi [][]complex128) [][]float64 {
	abs := make([][]float64, len(csi))
	for i := range csi {
		abs[i] = make([]float64, len(csi[i]))

		for j, value := range csi[i] {
			abs[i][j] = cmplx.Abs(value)
		}
	}

	return abs
}


func CsiToPhase(csi [][]complex128) [][]float64 {
	phase := make([][]float64, len(csi))
	for i := range csi {
		phase[i] = make([]float64, len(csi[i]))

		for j, value := range csi[i] {
			phase[i][j] = cmplx.Phase(value)
		}
	}

	return phase
}
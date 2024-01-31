package processor

import "math/cmplx"

func ReHandler(value complex128) float64 {
	return real(value)
}

func ImHandler(value complex128) float64 {
	return imag(value)
}

func PhaseHandler(value complex128) float64 {
	return cmplx.Phase(value)
}

func AbsHandler(value complex128) float64 {
	return cmplx.Abs(value)
}

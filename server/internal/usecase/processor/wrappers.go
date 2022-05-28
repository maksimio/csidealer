package processor


func reWrapper(value complex128) float64 {
	return real(value)
}

func imWrapper(value complex128) float64 {
	return imag(value)
}
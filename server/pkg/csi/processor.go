package csi

func CsiMap(csi [][]complex128, f func(complex128) float64) [][]float64 {
	data := make([][]float64, len(csi))
	for i := range csi {
		data[i] = make([]float64, len(csi[i]))

		for j, value := range csi[i] {
			data[i][j] = f(value)
		}
	}

	return data
}

package processor

import (
	"csidealer/internal/models"
	"math"
	"testing"
)

const EPS = 1e-10

func TestHandlers(t *testing.T) {
	testTable := []struct {
		csiValue complex128
		expected float64
		handler  func(complex128) float64
	}{
		{csiValue: 2 + 4i, expected: 2, handler: ReHandler},
		{csiValue: 2 + 4i, expected: 4, handler: ImHandler},
		{csiValue: 8 + 6i, expected: 10, handler: AbsHandler},
		{csiValue: 1 + 1i, expected: math.Pi / 4, handler: PhaseHandler},
		// важно, чтобы интервал PhaseHandler был [-Pi; Pi], а не [0; 2Pi]
		{csiValue: 2 - 2i, expected: -math.Pi / 4, handler: PhaseHandler},
	}

	mockProcessor := NewProcessorService(10)

	for _, testCase := range testTable {
		res := mockProcessor.CsiMap(models.Csi{[]complex128{testCase.csiValue}}, testCase.handler)
		if math.Abs(res[0][0]-testCase.expected) >= EPS {
			t.Log(res)
			t.Log(math.Abs(res[0][0] - testCase.expected))
			t.Errorf("некорректный результат: %f (исходное значение CSI: %f)", res[0][0], testCase.csiValue)
		}
	}
}

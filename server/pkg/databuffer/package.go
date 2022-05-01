package databuffer

import (
	"csidealer/pkg/csi"
	"github.com/google/uuid"
	"math/cmplx"
	"time"
)

type Package struct {
	Timestamp int64           `json:"ts"`
	Uuid      string          `json:"id"`
	Number    uint64          `json:"num"`
	Info      csi.PackageInfo `json:"info"`
	Data      [][]float64     `json:"data"`
}

func NewPackage(info csi.PackageInfo, data [][]float64, number uint64) *Package {
	p := new(Package)
	p.Info = info
	p.Data = data
	p.Number = number
	p.Uuid = uuid.New().String()
	p.Timestamp = time.Now().UnixMilli()
	return p
}

type PackageUnion struct {
	Abs   Package
	Phase Package
	Re    Package
	Im    Package
}

func NewPackageUnion(singlePackage csi.CsiPackage, number uint64) *PackageUnion {
	p := new(PackageUnion)

	abs := csi.CsiMap(singlePackage.Csi, cmplx.Abs)
	p.Abs = *NewPackage(singlePackage.PackageInfo, abs, number)

	phase := csi.CsiMap(singlePackage.Csi, cmplx.Phase)
	p.Phase = *NewPackage(singlePackage.PackageInfo, phase, number)

	re := csi.CsiMap(singlePackage.Csi, realWrapper)
	p.Re = *NewPackage(singlePackage.PackageInfo, re, number)

	im := csi.CsiMap(singlePackage.Csi, imagWrapper)
	p.Im = *NewPackage(singlePackage.PackageInfo, im, number)

	return p
}

func realWrapper(c complex128) float64 {
	return real(c)
}

func imagWrapper(c complex128) float64 {
	return imag(c)
}

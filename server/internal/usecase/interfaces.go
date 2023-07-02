package usecase

import "csidealer/internal/entity"

type (
	CsiUC interface {
		MoveRawTraffic(data []byte)
		FlushBuffer()
		GetTcpRemoteAddr() string
		SetTcpRemoteAddr(addr string)

		StartLog(filepath string) error
		StopLog() error
		GetLogWriteByteCount() uint64
		GetLogStartTime() int64
		IsLog() bool
		GetLogPackageCount() uint64

		GetCsi(csiType uint8, count int) ([]entity.ApiPackage, error)
		GetSubcarrier(csiType uint8, count, h, i int) ([]float64, error)
		GetCsiPackageCount() uint64
		GetCsiPackageMaxCount() uint64

		RoutersInfo() []entity.ApiRouterInfo
		RouterConnect(id, addr string) error
		RouterDisconnect(id string) error
		RouterSendDataRun(id, dstMacAddr string, numOfPacketToSend, pktIntervalUs, pktLen uint16)
		RouterSendDataStop(id string) error
		RouterClientMainRun(id, ip string, port int)
		RouterClientMainStop(id string) error

		GetPackageFilterLimits() (isActive bool, payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
		SetPackageFilterLimits(isActive bool, payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)

		OnPushPacket(cb func(entity.ApiPackageAbsPhase))
	}

	IRepo interface {
		Push(csiPackage *entity.Package)
		GetLastN(n int) []*entity.Package
		GetFullCount() uint64
		GetMaxCount() uint64
	}

	IBuffer interface {
		Push(data []byte)
		GetAllSplitted() []entity.RawPackage
		Flush()
	}

	IFSLogger interface {
		Start(filename string) error
		Stop()
		Write(data []byte) error
		IsOpen() bool
		GetStartTime() int64
		GetWriteByteCount() uint64
	}

	IProcessor interface {
		CsiMap(csi entity.Csi, f func(complex128) float64) [][]float64
		PackageMap(data []*entity.Package, handler func(complex128) float64) []entity.ApiPackage
		SubcarrierMap(data []*entity.Package, handler func(complex128) float64, h, i int) ([]float64, error)
	}

	IFilter interface {
		Check(info *entity.PackageInfo) bool
		GetLimits() (payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
		SetLimits(payloadLenMin, payloadLenMax uint16, nr, nc, nTones uint8)
	}

	IDecoder interface {
		DecodeCsiPackage([]byte) *entity.Package
	}

	IFSReader interface {
		List() ([]string, error)
		Start(filename string) error
		Stop() error
		GetDataPackage() []byte
		GetReadPercent() float64
		IsOpen() bool
	}

	IAtherosClient interface {
		GetId() string
		GetAddr() string
		Connect(addr string) error
		GetIsConnected() bool
		GetIsAvailable() bool
		Disconnect() error

		ClientMainRun(serverIP string, serverPort int) error
		GetIsClientMainActive() bool
		ClientMainStop() error

		SendDataRun(ifName, dstMacAddr string, numOfPacketToSend, pktIntervalUs, pktLen uint16) error
		GetIsSendDataActive() bool
		SendDataStop() error
	}
)

package config

// Конфиг сделано одноуровневым для метода Update
type Config struct {
	// IP-адреса и порты
	TcpPort  int    `yaml:"tcpPort"`
	RxIp     string `yaml:"rxIp"`
	TxIp     string `yaml:"txIp"`
	TargetIp string `yaml:"targetIp"` // Данный сервер, на который будет отправлена CSI

	// для функции передачи CSI
	IfName            string `yaml:"ifName"`
	DstMacAddr        string `yaml:"dstMacAddr"`
	NumOfPacketToSend int    `yaml:"numOfPacketToSend"`
	PktIntervalUs     int    `yaml:"pktIntervalUs"`
	PktLen            int    `yaml:"pktLen"`

	// различные параметры
	FilterOrder          int    `yaml:"filterOrder"`
	CsiLocalRepoMaxCount uint64 `yaml:"csiLocalRepoMaxCount"`
	DatFilePath          string `yaml:"datFilePath"`
	ProcessorRounder     int    `yaml:"ProcessorRounder"`

	Filter struct {
		PayloadLen struct {
			Min uint16 `yaml:"min"`
			Max uint16 `yaml:"max"`
		} `yaml:"payloadLen"`
		Nr     int `yaml:"nr"`
		Nc     int `yaml:"nc"`
		NTones int `yaml:"nTones"`
	} `yaml:"filter"`
}

package config

// Конфиг сделано одноуровневым для метода Update
type Config struct {
	// IP-адреса и порты
	HttpStaticPath string `yaml:"httpStaticPath"`
	HttpPort       int    `yaml:"httpPort"`
	TcpPort        int    `yaml:"tcpPort"`
	WebsocketPort  int    `yaml:"websocketPort"`
	RxIp           string `yaml:"rxIp"`
	TxIp           string `yaml:"txIp"`
	TargetIp       string `yaml:"targetIp"` // Данный сервер, на который будет отправлена CSI

	// для функции передачи CSI
	IfName            string `yaml:"ifName"`
	DstMacAddr        string `yaml:"dstMacAddr"`
	NumOfPacketToSend uint16 `yaml:"numOfPacketToSend"`
	PktIntervalUs     uint16 `yaml:"pktIntervalUs"`
	PktLen            uint16 `yaml:"pktLen"`

	// различные параметры
	SmoothOrder          int    `yaml:"smoothOrder"`
	CsiLocalRepoMaxCount uint64 `yaml:"csiLocalRepoMaxCount"`
	DatFilePath          string `yaml:"datFilePath"`
	ProcessorRounder     int    `yaml:"ProcessorRounder"`

	Filter struct {
		PayloadLen struct {
			Min uint16 `yaml:"min"`
			Max uint16 `yaml:"max"`
		} `yaml:"payloadLen"`
		Nr     uint8 `yaml:"nr"`
		Nc     uint8 `yaml:"nc"`
		NTones uint8 `yaml:"nTones"`
	} `yaml:"filter"`
}

package config

// Конфиг сделано одноуровневым для метода Update
type FileScheme struct {
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
	FilterOrder int `yaml:"filterOrder"`
}

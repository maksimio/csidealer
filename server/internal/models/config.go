package models

type Config struct {
	// Роутеры
	Rx       RouterConfigInfo `yaml:"rx"`
	Tx       RouterConfigInfo `yaml:"tx"`
	SendData SendDataInfo     `yaml:"sendData"`

	// IP-адреса и порты
	HttpStaticPath string `yaml:"httpStaticPath"`
	HttpPort       int    `yaml:"httpPort"`
	TcpPort        int    `yaml:"tcpPort"`
	WebsocketPort  int    `yaml:"websocketPort"`
	RxIp           string `yaml:"rxIp"`
	TxIp           string `yaml:"txIp"`
	TargetIp       string `yaml:"targetIp"` // Данный сервер, на который будет отправлена CSI

	// для функции передачи CSI

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

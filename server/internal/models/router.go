package models

type ApiRouterInfo struct {
	Id                 string `json:"id"`
	Addr               string `json:"addr"`
	IsConnected        bool   `json:"is_connected"`
	IsClientMainActive bool   `json:"is_clientmain_active"`
	IsSendDataActive   bool   `json:"is_sendData_active"`
	IsAvailable        bool   `json:"is_available"`
}

type RouterInfo struct {
	Username string `yaml:"username"`
	IpAddr   string `yaml:"ipAddr"`
}

type SendDataInfo struct {
	IfName            string `yaml:"ifName"`
	DstMacAddr        string `yaml:"dstMacAddr"`
	NumOfPacketToSend uint16 `yaml:"numOfPacketToSend"`
	PktIntervalUs     uint16 `yaml:"pktIntervalUs"`
	PktLen            uint16 `yaml:"pktLen"`
}

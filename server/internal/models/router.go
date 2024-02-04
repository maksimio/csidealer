package models

type ApiRoutersStatus struct {
	Rx         ApiRouterInfo `json:"rx"`
	Tx         ApiRouterInfo `json:"tx"`
	SendData   SendDataInfo  `json:"sendData"`
	ServerIp   string        `json:"serverIp"`
	ServerPort int           `json:"serverPort"`
}

type ApiRouterInfo struct {
	Id                 string `json:"id"`
	Addr               string `json:"addr"`
	IsConnected        bool   `json:"is_connected"`
	IsClientMainActive bool   `json:"is_clientmain_active"`
	IsSendDataActive   bool   `json:"is_sendData_active"`
}

type RouterConfigInfo struct {
	Username string `yaml:"username"`
	IpAddr   string `yaml:"ipAddr"`
}

type SendDataInfo struct {
	IfName            string `yaml:"ifName" json:"ifName"`
	DstMacAddr        string `yaml:"dstMacAddr" json:"dstMacAddr"`
	NumOfPacketToSend uint16 `yaml:"numOfPacketToSend" json:"numOfPacketToSend"`
	PktIntervalUs     uint16 `yaml:"pktIntervalUs" json:"pktIntervalUs"`
	PktLen            uint16 `yaml:"pktLen" json:"pktLen"`
}

package entity

type ApiRouterInfo struct {
	Id                 string `json:"id"`
	Addr               string `json:"addr"`
	IsConnected        bool   `json:"is_connected"`
	IsClientMainActive bool   `json:"is_clientmain_active"`
	IsSendDataActive   bool   `json:"is_sendData_active"`
	IsAvailable        bool   `json:"is_available"`
}

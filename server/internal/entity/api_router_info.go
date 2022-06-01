package entity

type ApiRouterInfo struct {
	Id                 string `json:"id"`
	Addr               string `json:"addr"`
	IsConnected        bool   `json:"isConnected"`
	IsClientMainActive bool   `json:"isClientMainActive"`
	IsSendDataActive   bool   `json:"isSendDataActive"`
}

package main

import (
	"csidealer/internal/apiserver"
	"csidealer/internal/tcpserver"
	"csidealer/pkg/csi"
)

func main() {
	c := make(chan csi.CsiPackage)
	go tcpserver.RunTcpServer(8081, c)
	apiserver.RunApiServer(80, c, "../client/build")
}

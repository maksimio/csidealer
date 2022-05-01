package main

import (
	"csidealer/internal/apiserver"
	"csidealer/internal/tcpserver"
	"csidealer/pkg/csicore"
)

func main() {
	c := make(chan *csicore.CsiPackage)
	go tcpserver.RunTcpServer(8081, c)
	apiserver.RunApiServer(80, c, "../client/build")
}
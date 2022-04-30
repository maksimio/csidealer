package main

import (
	"csidealer/internal/apiserver"
	"csidealer/internal/tcpserver"
)

func main() {
	go tcpserver.RunTcpServer(8081)
	apiserver.RunApiServer(80, "../client/build")
}
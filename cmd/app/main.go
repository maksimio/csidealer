package main

import "fmt"
import "csi_dealer/internal/tcpserver"

func main() {
	fmt.Println("Hello, world!")
	tcpserver.StartServer(8081)
}

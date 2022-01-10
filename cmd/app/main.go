package main

import "fmt"
import "csidealer/internal/tcpserver"

func main() {
	fmt.Println("Hello, world!")
	tcpserver.StartServer(8081)
}


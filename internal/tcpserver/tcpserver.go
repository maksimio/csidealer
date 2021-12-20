package tcpserver

import (
	"csi_dealer/pkg/databuffer"
	"fmt"
	"net"
	"os"
)

const (
	BUF_LEN = 2048
)

func StartServer(port int) {
	ln, err := net.Listen("tcp", ":"+fmt.Sprint(port))
	if err != nil {
		fmt.Println("Ошибка прослушки на порту", port, ":", err)
		os.Exit(1)
	}

	defer ln.Close()

	for {
		fmt.Println("Сервер ожидает подключение на", port, "порту")
		conn, _ := ln.Accept()
		trafficBuf := databuffer.NewTrafficBuffer()
		readLoop(conn, trafficBuf)
	}
}

func readLoop(conn net.Conn, trafficBuf *databuffer.TrafficBuffer) {
	fmt.Println("Новое подключение от:", conn.RemoteAddr())
	buf := make([]byte, BUF_LEN)

	for {
		readCount, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Ошибка чтения из сокета:", err, "Отключение сервера")
			break
		}
		// x := binary.LittleEndian.Uint32(buf)
		trafficBuf.Push(buf[:readCount])
		fmt.Println("Length:", trafficBuf.Length())
	}
}

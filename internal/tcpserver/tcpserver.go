package tcpserver

import (
	"csidealer/pkg/databuffer"
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
		buf := databuffer.NewBufferFlow()
		readLoop(conn, buf)
	}
}

func readLoop(conn net.Conn, buf *databuffer.BufferFlow) {
	fmt.Println("Новое подключение от:", conn.RemoteAddr())
	data := make([]byte, BUF_LEN)

	for {
		readCount, err := conn.Read(data)
		if err != nil {
			fmt.Println("Ошибка чтения из сокета:", err, "Отключение сервера")
			break
		}
		fmt.Println("Принято:", readCount)
		buf.Push(data[:readCount])
	}
}

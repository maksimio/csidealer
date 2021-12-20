package tcpserver

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	BUF_LEN = 4
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
		readLoop(conn)
	}
}

func readLoop(conn net.Conn) {
	fmt.Println("Новое подключение от:", conn.RemoteAddr())
	buf := make([]byte, BUF_LEN)

	for {
		readCount, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Ошибка чтения из сокета:", err, "Отключение сервера")
			break
		}
		fmt.Println(readCount, buf)
		// dec := make([]byte, 4)
		// l, _ := base64.RawStdEncoding.Decode(dec, buf)
		x := binary.LittleEndian.Uint32(buf)
		fmt.Println(x)
	}
}

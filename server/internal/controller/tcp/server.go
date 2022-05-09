package tcp

import (
	"csidealer/internal/usecase"
	"fmt"
	"net"
	"os"
)

const _buf_len = 2048

type TcpServer struct {
	csiUc usecase.Csi
	port  string
}

func NewTcpServer(uc usecase.Csi, port int) *TcpServer {
	return &TcpServer{
		csiUc: uc,
		port:  ":" + fmt.Sprint(port),
	}
}

func (s *TcpServer) Run() {
	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		fmt.Println("Ошибка прослушки на порту", s.port, ":", err)
		os.Exit(1) //TODO сделать более элегантно
	}

	defer ln.Close()

	for {
		fmt.Println("Сервер ожидает подключение на", s.port, "порту")
		conn, _ := ln.Accept()
		s.listenConnection(conn)
	}
}

func (s *TcpServer) listenConnection(conn net.Conn) {
	fmt.Println("Новое подключение от:", conn.RemoteAddr())
	data := make([]byte, _buf_len)

	for {
		readCount, err := conn.Read(data)
		if err != nil {
			fmt.Println("Ошибка чтения из сокета:", err)
			break
		}

		s.csiUc.HandleRawTraffic(data[:readCount])
	}
}

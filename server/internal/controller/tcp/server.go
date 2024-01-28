package tcp

import (
	"csidealer/internal/usecase"
	"fmt"
	"log"
	"net"
)

const _buf_len = 2048

type TcpServer struct {
	csiUc usecase.CsiUC
	port  string
}

func NewTcpServer(uc usecase.CsiUC, port int) *TcpServer {
	return &TcpServer{
		csiUc: uc,
		port:  ":" + fmt.Sprint(port),
	}
}

func (s *TcpServer) Run() {
	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatal("ошибка открытия порта", s.port, ":", err)
	}

	defer ln.Close()

	for {
		log.Printf("TCP-сервер ожидает подключение на %s порту", s.port)
		conn, _ := ln.Accept()
		s.listenConnection(conn)
	}
}

func (s *TcpServer) listenConnection(conn net.Conn) {
	s.csiUc.FlushBuffer()
	log.Printf("новое подключение от %s", conn.RemoteAddr())
	s.csiUc.SetTcpRemoteAddr(conn.RemoteAddr().String())
	data := make([]byte, _buf_len)

	for {
		readCount, err := conn.Read(data)
		if err != nil {
			log.Print("ошибка чтения из сокета:", err)
			s.csiUc.SetTcpRemoteAddr("")
			break
		}

		s.csiUc.MoveRawTraffic(data[:readCount])
	}
}

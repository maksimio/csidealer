package tcp

import (
	"csidealer/internal/services/buffer"
	"fmt"
	"log"
	"net"
)

const _buf_len = 2048

type TcpController struct {
	bufferService *buffer.BufferService
	port          string
}

func NewTcpController(bufferService *buffer.BufferService, port int) *TcpController {
	return &TcpController{
		bufferService: bufferService,
		port:          ":" + fmt.Sprint(port),
	}
}

func (s *TcpController) Run() {
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

func (s *TcpController) listenConnection(conn net.Conn) {
	s.bufferService.Flush()
	log.Printf("новое подключение от %s", conn.RemoteAddr())
	s.bufferService.TcpRemoteAddr = conn.RemoteAddr().String()
	data := make([]byte, _buf_len)

	for {
		readCount, err := conn.Read(data)
		if err != nil {
			log.Print("соединение разорвано:", err)
			s.bufferService.TcpRemoteAddr = ""
			break
		}
		s.bufferService.Push(data[:readCount])
	}
}

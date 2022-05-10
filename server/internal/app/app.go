package app

import (
	"csidealer/internal/controller/http"
	"csidealer/internal/controller/tcp"
	"csidealer/internal/controller/websocket"
	"csidealer/internal/usecase"
	"csidealer/internal/usecase/buffer"
	"csidealer/internal/usecase/file_writer"
	"csidealer/internal/usecase/repo"
)

func Run() {
	csiUseCase := usecase.NewCsiUseCase(
		repo.NewCsiLocalRepo(20),
		buffer.NewCsiRawRepo(),
		file_writer.NewFileWriter(),
	)

	tcpServer := tcp.NewTcpServer(csiUseCase, 8081)
	websocketServer := websocket.NewWebsocketServer(csiUseCase, 7000)
	httpServer := http.NewHttpServer(csiUseCase, 80, "./build")

	go tcpServer.Run()
	go websocketServer.Run()
	httpServer.Run()
}

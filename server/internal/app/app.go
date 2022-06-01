package app

import (
	"csidealer/internal/controller/http"
	"csidealer/internal/controller/tcp"
	"csidealer/internal/controller/websocket"
	"csidealer/internal/usecase"
	"csidealer/internal/usecase/buffer"
	"csidealer/internal/usecase/decoder"
	"csidealer/internal/usecase/filter"
	"csidealer/internal/usecase/fs_logger"
	"csidealer/internal/usecase/processor"
	"csidealer/internal/usecase/repo"
)

func Run() {
	csiUseCase := usecase.NewCsiUseCase(
		repo.NewCsiLocalRepo(1000),
		buffer.NewCsiRawRepo(),
		fs_logger.NewFileLogger("./logs/"),
		processor.NewProcessor(3),
		filter.NewFilter(500, 1800, 2, 2, 56),
		decoder.NewCsiDecoder(),
	)

	return


	tcpServer := tcp.NewTcpServer(csiUseCase, 8081)
	websocketServer := websocket.NewWebsocketServer(csiUseCase, 7000)
	httpServer := http.NewHttpServer(csiUseCase, 80, "./build")

	go tcpServer.Run()
	go websocketServer.Run()
	httpServer.Run()
}

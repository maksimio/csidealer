package app

import (
	"csidealer/config"
	"csidealer/internal/controllers/http"
	"csidealer/internal/controllers/tcp"
	"csidealer/internal/controllers/websocket"
	"csidealer/internal/usecase"
	"csidealer/internal/usecase/buffer"
	"csidealer/internal/usecase/decoder"
	"csidealer/internal/usecase/filter"
	"csidealer/internal/usecase/fs_logger"
	"csidealer/internal/usecase/processor"
	"csidealer/internal/usecase/repo"
	"csidealer/internal/usecase/ssh"
	"log"
)

func Run() {
	config, _ := config.ReadConfig()

	clients := []*ssh.AtherosClient{
		ssh.NewAtherosClient("root"),
		ssh.NewAtherosClient("root"),
	}

	routers := make([]*usecase.IAtherosClient, len(clients))
	for i, v := range clients {
		iRouter := usecase.IAtherosClient(v)
		routers[i] = &iRouter // TODO: Разобраться, как лучше работать с указателями
	}

	csiUseCase := usecase.NewCsiUseCase(
		repo.NewCsiLocalRepo(config.CsiLocalRepoMaxCount),
		buffer.NewCsiRawRepo(),
		fs_logger.NewFileLogger(config.DatFilePath),
		processor.NewProcessor(config.ProcessorRounder),
		filter.NewFilter(
			config.Filter.PayloadLen.Min,
			config.Filter.PayloadLen.Max,
			config.Filter.Nr,
			config.Filter.Nc,
			config.Filter.NTones,
		),
		decoder.NewCsiDecoder(),
		routers,
		config.SmoothOrder,
	)

	tcpServer := tcp.NewTcpServer(csiUseCase, config.TcpPort)
	websocketServer := websocket.NewWebsocketServer(csiUseCase, config.WebsocketPort)
	httpServer := http.NewHttpServer(csiUseCase, config.HttpPort, config.HttpStaticPath)

	go tcpServer.Run()
	go httpServer.Run()

	log.Print("запуск передачи пакетов")
	rx := *routers[0]
	tx := *routers[1]
	rx.Connect(config.RxIp)
	tx.Connect(config.TxIp)
	rx.ClientMainRun(config.TargetIp, config.TcpPort)
	tx.SendDataRun(
		config.IfName,
		config.DstMacAddr,
		config.NumOfPacketToSend,
		config.PktIntervalUs,
		config.PktLen,
	)

	log.Print("запуск сервера websocket...")
	websocketServer.Run()

}

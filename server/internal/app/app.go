package app

import (
	"csidealer/config"
	"csidealer/internal/controllers/http"
	"csidealer/internal/controllers/tcp"
	"csidealer/internal/controllers/websocket"
	"csidealer/internal/models"
	"csidealer/internal/services"
	"csidealer/internal/services/buffer"
	"csidealer/internal/services/raw_writer"
	"csidealer/internal/services/ssh"
	"log"
)

func Run() {
	config, _ := config.ReadConfig()

	clients := []*ssh.AtherosClient{
		ssh.NewAtherosClient("root"),
		ssh.NewAtherosClient("root"),
	}

	routers := make([]*services.IAtherosClient, len(clients))
	for i, v := range clients {
		iRouter := services.IAtherosClient(v)
		routers[i] = &iRouter // TODO: Разобраться, как лучше работать с указателями
	}

	// csiUseCase := services.NewCsiUseCase(
	// 	repo.NewCsiLocalRepo(config.CsiLocalRepoMaxCount),
	// 	buffer.NewBufferService(),
	// 	fs_logger.NewFileLogger(config.DatFilePath),
	// 	processor.NewProcessor(config.ProcessorRounder),
	// 	filter.NewFilter(
	// 		config.Filter.PayloadLen.Min,
	// 		config.Filter.PayloadLen.Max,
	// 		config.Filter.Nr,
	// 		config.Filter.Nc,
	// 		config.Filter.NTones,
	// 	),
	// 	decoder.NewCsiDecoder(),
	// 	routers,
	// 	config.SmoothOrder,
	// )

	rawData := make(chan models.RawPackage)

	bufferService := buffer.NewBufferService(rawData)
	rawWriterService := raw_writer.NewRawWriterService(rawData, config.DatFilePath)
	go rawWriterService.Log()

	tcpServer := tcp.NewTcpServer(bufferService, config.TcpPort)

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

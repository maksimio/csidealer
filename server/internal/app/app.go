package app

import (
	"csidealer/config"
	"csidealer/internal/controllers/http"
	"csidealer/internal/controllers/tcp"
	"csidealer/internal/controllers/websocket"

	// "csidealer/internal/controllers/websocket"
	"csidealer/internal/models"
	// "csidealer/internal/services"
	"csidealer/internal/services/buffer"
	"csidealer/internal/services/decoder"
	"csidealer/internal/services/filter"
	"csidealer/internal/services/processor"
	"csidealer/internal/services/raw_writer"
	"csidealer/internal/services/storage"
	// "log"
)

func Run() {
	config, _ := config.ReadConfig()

	// routers := []*ssh_client.SshClient{
	// 	ssh_client.NewSshClient("root"),
	// 	ssh_client.NewSshClient("root"),
	// }

	// csiUseCase := services.NewCsiUseCase(
	// 	repo.NewCsiLocalRepo(config.CsiLocalRepoMaxCount),
	// 	buffer.NewBufferService(),
	// 	fs_logger.NewFileLogger(config.DatFilePath),
	// 	processor.NewProcessor(config.ProcessorRounder),
	// 	filter.NewFilter(
	// config.Filter.PayloadLen.Min,
	// config.Filter.PayloadLen.Max,
	// config.Filter.Nr,
	// config.Filter.Nc,
	// config.Filter.NTones,
	// 	),
	// 	decoder.NewCsiDecoder(),
	// 	routers,
	// 	config.SmoothOrder,
	// )

	toRawWriter := make(chan models.RawPackage)
	toDecoder := make(chan models.RawPackage)
	toFilter := make(chan models.Package)
	toStorage := make(chan models.Package)
	toWebsocket := make(chan models.Package)

	bufferService := buffer.NewBufferService([]chan<- models.RawPackage{toRawWriter, toDecoder})
	rawWriterService := raw_writer.NewRawWriterService(toRawWriter, config.DatFilePath)
	decoderService := decoder.NewDecoderService(toDecoder, []chan<- models.Package{toFilter})
	filterService := filter.NewFilterService(
		toFilter,
		[]chan<- models.Package{toStorage, toWebsocket},
		config.Filter.PayloadLen.Min,
		config.Filter.PayloadLen.Max,
		config.Filter.Nr,
		config.Filter.Nc,
		config.Filter.NTones,
	)
	storageService := storage.NewStorageService(toStorage, config.CsiLocalRepoMaxCount)
	processorService := processor.NewProcessorService(config.ProcessorRounder)

	tcpController := tcp.NewTcpController(bufferService, config.TcpPort)
	httpController := http.NewHttpController(bufferService, rawWriterService, config.HttpPort, config.HttpStaticPath)
	websocketController := websocket.NewWebsocketController(toWebsocket, processorService, config.WebsocketPort)

	go rawWriterService.Run()
	go decoderService.Run()
	go filterService.Run()
	go storageService.Run()

	go tcpController.Run()
	go httpController.Run()

	// -------------------------------------
	// log.Print("запуск передачи пакетов")
	// rx := *routers[0]
	// tx := *routers[1]
	// rx.Connect(config.RxIp)
	// tx.Connect(config.TxIp)
	// rx.ClientMainRun(config.TargetIp, config.TcpPort)
	// tx.SendDataRun(
	// 	config.IfName,
	// 	config.DstMacAddr,
	// 	config.NumOfPacketToSend,
	// 	config.PktIntervalUs,
	// 	config.PktLen,
	// )

	// log.Print("запуск сервера websocket...")
	// -------------------------------------

	websocketController.Run()

}

package app

import (
	"csidealer/config" // TODO: в internal не должно быть внешних зависимостей
	"csidealer/internal/controllers/http"
	"csidealer/internal/controllers/tcp"
	"csidealer/internal/controllers/websocket"
	"csidealer/internal/models"
	"csidealer/internal/services/buffer"
	"csidealer/internal/services/decoder"
	"csidealer/internal/services/filter"
	"csidealer/internal/services/processor"
	"csidealer/internal/services/raw_writer"
	"csidealer/internal/services/storage"
)

func Run(conf config.Config) {

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
		conf.Filter.PayloadLen.Min,
		conf.Filter.PayloadLen.Max,
		conf.Filter.Nr,
		conf.Filter.Nc,
		conf.Filter.NTones,
	)
	storageService := storage.NewStorageService(toStorage, conf.CsiLocalRepoMaxCount)
	processorService := processor.NewProcessorService(conf.ProcessorRounder)

	tcpController := tcp.NewTcpController(bufferService, conf.TcpPort)
	httpController := http.NewHttpController(bufferService, rawWriterService, conf.HttpPort, conf.HttpStaticPath)
	websocketController := websocket.NewWebsocketController(toWebsocket, processorService, conf.WebsocketPort)

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
	// -------------------------------------

	websocketController.Run()

}

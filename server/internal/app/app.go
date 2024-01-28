package app

import (
	"csidealer/config"
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
	"csidealer/internal/usecase/ssh"
	"log"
)

func Run() {
	conf, err := config.ReadConfig()
	if err != nil {
		log.Print("возникла ошибка при чтении конфигурационного файла: ", err)
		return
	} else {
		log.Print("конфигурационный файл успешно прочитан")
	}

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
		repo.NewCsiLocalRepo(1000),
		buffer.NewCsiRawRepo(),
		fs_logger.NewFileLogger("./logs/"),
		processor.NewProcessor(3),
		filter.NewFilter(500, 1800, 2, 2, 56),
		decoder.NewCsiDecoder(),
		routers,
	)

	tcpServer := tcp.NewTcpServer(csiUseCase, conf.Tcp.Port)
	websocketServer := websocket.NewWebsocketServer(csiUseCase, 8082)
	httpServer := http.NewHttpServer(csiUseCase, 80, "./build")

	go tcpServer.Run()
	go httpServer.Run()

	rx := *routers[0]
	tx := *routers[1]
	rx.Connect(conf.Rx.Ip)
	tx.Connect(conf.Tx.Ip)
	rx.ClientMainRun(conf.Rx.TargetIp, conf.Tcp.Port)
	tx.SendDataRun(conf.Tx.IfName, conf.Tx.DstMacAddr, uint16(conf.Tx.NumOfPacketToSend), uint16(conf.Tx.PktIntervalUs), uint16(conf.Tx.PktLen))

	websocketServer.Run()

}

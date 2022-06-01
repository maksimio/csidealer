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
	"csidealer/internal/usecase/ssh"
)

func Run() {
	clients := []*ssh.AtherosClient{
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

	tcpServer := tcp.NewTcpServer(csiUseCase, 8081)
	websocketServer := websocket.NewWebsocketServer(csiUseCase, 7000)
	httpServer := http.NewHttpServer(csiUseCase, 80, "./build")

	go tcpServer.Run()
	go websocketServer.Run()

	// cl1 := ssh.NewAtherosClient("root")
	// if err := cl1.Connect("192.168.1.100"); err != nil {
	// 	fmt.Println(err)
	// }

	// cl2 := ssh.NewAtherosClient("root")
	// if err := cl2.Connect("192.168.1.1"); err != nil {
	// 	fmt.Println(err)
	// }

	// time.Sleep(5 * time.Second)
	// cl2.ClientMainStop()
	// cl1.SendDataStop()
	// time.Sleep(1 * time.Second)
	// go cl2.ClientMainStart("192.168.1.231", "8081")
	// go cl1.SendDataRun("wlan0", "00:7F:5D:3E:4A", 100, 50, 1000)
	// time.Sleep(7 * time.Second)
	// if err := cl1.SendDataStop(); err != nil {
	// 	fmt.Println(err)
	// }

	httpServer.Run()
}

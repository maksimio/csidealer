package app

import (
	"csidealer/internal/controller/tcp"
	"csidealer/internal/usecase"
	"csidealer/internal/usecase/buffer"
	"csidealer/internal/usecase/file_writer"
	"csidealer/internal/usecase/repo"
)

func Run() {
	csiUseCase := usecase.NewCsiUseCase(
		&repo.CsiLocalRepo{},
		buffer.NewCsiRawRepo(),
		file_writer.NewFileWriter(),
	)

	tcpServer := tcp.NewTcpServer(csiUseCase, 8081)

	tcpServer.Run()
}

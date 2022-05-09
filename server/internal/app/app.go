package app

import (
	"csidealer/internal/usecase"
	"csidealer/internal/usecase/buffer"
	"csidealer/internal/usecase/file_writer"
	"csidealer/internal/usecase/repo"
	"fmt"
)

func Run() {
	csiUseCase := usecase.NewCsiUseCase(
		&repo.CsiLocalRepo{},
		buffer.NewCsiRawRepo(),
		file_writer.NewFileWriter(),
	)

	fmt.Println(csiUseCase)
}

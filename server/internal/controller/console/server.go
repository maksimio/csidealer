package console

import "csidealer/internal/usecase"

type ConsoleServer struct {
}

func NewConsoleServer(uc usecase.CsiUC) *ConsoleServer {
	return &ConsoleServer{}
}

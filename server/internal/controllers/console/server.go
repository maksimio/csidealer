package console

import "csidealer/internal/services"

type ConsoleServer struct {
}

func NewConsoleServer(uc services.CsiUC) *ConsoleServer {
	return &ConsoleServer{}
}

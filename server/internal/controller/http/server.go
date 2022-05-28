package http

import (
	"csidealer/internal/usecase"
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	port   string
	router *gin.Engine
	uiPath string
	api    Api
}

func NewHttpServer(uc usecase.CsiUC, port int, uiPath string) *HttpServer {
	router := gin.Default()
	routGr := router.Group("/api/v1")

	return &HttpServer{
		port:   "localhost:" + fmt.Sprint(port),
		uiPath: uiPath,
		api:    NewApiV1(routGr, uc),
		router: router,
	}
}

func (s HttpServer) Run() {
	s.api.Register()

	s.router.Use(static.Serve("/", static.LocalFile(s.uiPath, true)))
	fmt.Println("HTTP-сервер ожидает подключение на", s.port, "порту")
	s.router.Run(s.port)
}

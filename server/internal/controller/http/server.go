package http

import (
	"csidealer/internal/usecase"
	"fmt"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	port   string
	router *gin.Engine
}

func NewHttpServer(uc usecase.CsiUC, port int, uiPath string) *HttpServer {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(uiPath, true)))
	router.Use(cors.Default())

	routGr := router.Group("/api/v1")
	api := NewApiV1(routGr, uc)
	api.Register()

	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	return &HttpServer{
		port:   "localhost:" + fmt.Sprint(port),
		router: router,
	}
}

func (s HttpServer) Run() {
	fmt.Println("HTTP-сервер ожидает подключение на", s.port, "порту")
	s.router.Run(s.port)
}

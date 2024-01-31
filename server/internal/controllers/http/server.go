package http

import (
	"csidealer/internal/services/buffer"
	"csidealer/internal/services/raw_writer"
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type HttpController struct {
	port   string
	router *gin.Engine
	api    *ApiV1
}

func NewHttpController(
	bufferService *buffer.BufferService,
	rawWriterService *raw_writer.RawWriterService,
	port int, uiPath string) *HttpController {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile(uiPath, true)))
	router.Use(cors.Default())

	routGr := router.Group("/api/v1")
	api := NewApiV1(bufferService, rawWriterService, routGr) // TODO: сервисы передать сюда
	api.Register()

	return &HttpController{
		port:   "localhost:" + fmt.Sprint(port),
		router: router,
		api:    api,
	}
}

func (s HttpController) Run() {
	log.Printf("HTTP-сервер ожидает подключение на %s порту", s.port)
	s.router.Run(s.port)
}

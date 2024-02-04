package http

import (
	"csidealer/internal/services/buffer"
	"csidealer/internal/services/raw_writer"
	"csidealer/internal/services/router_connector"
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type HttpController struct {
	port   string
	router *gin.Engine
	routGr *gin.RouterGroup

	bufferService          *buffer.BufferService
	rawWriterService       *raw_writer.RawWriterService
	routerConnectorService *router_connector.RouterConnectorService
}

func NewHttpController(
	bufferService *buffer.BufferService,
	rawWriterService *raw_writer.RawWriterService,
	routerConnectorService *router_connector.RouterConnectorService,
	port int, uiPath string) *HttpController {
	// --- ИНИЦИАЛЦИЗАЦИЯ ---
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(uiPath, true)))
	router.Use(cors.Default())
	routGr := router.Group("/api/v1")

	httpController := &HttpController{
		port:   "localhost:" + fmt.Sprint(port),
		router: router,
		routGr: routGr,

		bufferService:          bufferService,
		rawWriterService:       rawWriterService,
		routerConnectorService: routerConnectorService,
	}
	// --- МАРШРУТЫ ---
	// --- Запись сырых данных CSI
	log := routGr.Group("/write")
	log.GET("/start", httpController.startLog)
	log.GET("/stop", httpController.stopLog)
	log.GET("/state", httpController.stateLog)

	// --- Фильтрация данных

	// --- Команды роутерам
	routers := routGr.Group("/routers")
	routers.POST("/connect", httpController.connectRouters)
	routers.POST("/start", httpController.startCsiTransmit)
	routers.POST("/stop", httpController.stopCsiTransmit)
	routers.GET("/status", httpController.routersStatus)

	// --- Запрос и регулировка конфигурационных параметров

	return httpController
}

func (s HttpController) Run() {
	log.Printf("HTTP-сервер ожидает подключение на %s порту", s.port)
	s.router.Run(s.port)
}

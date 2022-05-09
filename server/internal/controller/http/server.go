package http

import (
	"csidealer/internal/usecase"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	csiUc  usecase.Csi
	port   string
	routGr *gin.RouterGroup
	uiPath string
}

func NewHttpServer(uc usecase.Csi, port int, uiPath string) *HttpServer {
	return &HttpServer{
		csiUc:  uc,
		port:   ":" + fmt.Sprint(port),
		uiPath: uiPath,
	}
}

func (s HttpServer) Run() {
	router := gin.Default()
	s.routGr = router.Group("/api/v1")

	s.registerApiV1()

	router.Use(static.Serve("/", static.LocalFile(s.uiPath, true)))
	fmt.Println("HTTP-сервер ожидает подключение на", s.port, "порту")
	router.Run(s.port)
}

func (s HttpServer) registerApiV1() {
	s.routGr.GET("/csiLastN/:type", s.csiLastN)
	s.routGr.GET("/subcarrierLastN/:type", s.subcarrierLastN)

	s.routGr.GET("/startLog", s.startLog)
	s.routGr.GET("/stopLog", s.stopLog)

	s.routGr.GET("/status", s.status)
	s.routGr.POST("/config", s.config)
}

func (s HttpServer) csiLastN(c *gin.Context) {
	// csiType := c.Param("type")
	// n, _ := strconv.Atoi(c.Query("n"))
	c.JSON(200, 1)
}

func (s HttpServer) subcarrierLastN(c *gin.Context) {
	// csiType := c.Param("type")
	// h, _ := strconv.Atoi(c.Query("h"))
	// index, _ := strconv.Atoi(c.Query("index"))
	// n, _ := strconv.Atoi(c.Query("n"))
	c.JSON(200, 1)
}

func (s HttpServer) startLog(c *gin.Context) {
	filepath := c.Query("filepath")
	fmt.Println(filepath)
	c.AbortWithStatus(200)
}

func (s HttpServer) stopLog(c *gin.Context) {
	fmt.Println("Стоп логирование")
	c.AbortWithStatus(200)
}

func (s HttpServer) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Информация о подключенном устройстве: статус подключения, IP, время подключения, число переданных пакетов",
	})
}

func (s HttpServer) config(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Конфигурация пути сохранения файла",
	})
}

package http

import (
	"csidealer/internal/usecase"
	// "fmt"
	"github.com/gin-gonic/gin"
)

type ApiV1 struct {
	routGr *gin.RouterGroup
	csiUc  usecase.Csi
}

func NewApiV1(rg *gin.RouterGroup, uc usecase.Csi) *ApiV1 {
	return &ApiV1{
		routGr: rg,
		csiUc:  uc,
	}
}

func (a *ApiV1) Register() {
	a.routGr.GET("/csiLastN/:type", a.csiLastN)
	a.routGr.GET("/subcarrierLastN/:type", a.subcarrierLastN)

	a.routGr.GET("/startLog", a.startLog)
	a.routGr.GET("/stopLog", a.stopLog)

	a.routGr.GET("/status", a.status)
	a.routGr.POST("/config", a.config)
}

func (a *ApiV1) csiLastN(c *gin.Context) {
	// csiType := c.Param("type")
	// n, _ := strconv.Atoi(c.Query("n"))
	c.JSON(200, 1)
}

func (a *ApiV1) subcarrierLastN(c *gin.Context) {
	// csiType := c.Param("type")
	// h, _ := strconv.Atoi(c.Query("h"))
	// index, _ := strconv.Atoi(c.Query("index"))
	// n, _ := strconv.Atoi(c.Query("n"))
	c.JSON(200, 1)
}

func (a *ApiV1) startLog(c *gin.Context) {
	filepath := c.Query("filepath")
	err := a.csiUc.StartLog(filepath)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (a *ApiV1) stopLog(c *gin.Context) {
	a.csiUc.StopLog()
	c.AbortWithStatus(200)
}

func (a *ApiV1) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Информация о подключенном устройстве: статус подключения, IP, время подключения, число переданных пакетов",
	})
}

func (a *ApiV1) config(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Конфигурация пути сохранения файла",
	})
}

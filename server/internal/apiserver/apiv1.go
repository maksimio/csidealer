package apiserver

import (
	"csidealer/pkg/databuffer"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiV1 struct {
	routGr *gin.RouterGroup
	buf    *databuffer.PackageBuffer
}

func NewApiV1(routGr *gin.RouterGroup, buf *databuffer.PackageBuffer) *ApiV1 {
	p := new(ApiV1)
	p.routGr = routGr
	p.buf = buf

	p.routGr.GET("/csiLastN/:type", p.csiLastN)
	p.routGr.GET("/subcarrierLastN/:type", p.subcarrierLastN)

	p.routGr.GET("/startLog", p.startLog)
	p.routGr.GET("/stopLog", p.stopLog)

	p.routGr.GET("/status", p.status)
	p.routGr.POST("/config", p.config)

	return p
}

func (api *ApiV1) csiLastN(c *gin.Context) {
	csiType := c.Param("type")
	n, _ := strconv.Atoi(c.Query("n"))
	data := api.buf.CsiLastN(n, csiType)
	c.JSON(200, data)
}

func (api *ApiV1) subcarrierLastN(c *gin.Context) {
	csiType := c.Param("type")
	h, _ := strconv.Atoi(c.Query("h"))
	index, _ := strconv.Atoi(c.Query("index"))
	n, _ := strconv.Atoi(c.Query("n"))
	data := api.buf.SubcarrierLastN(h, index, n, csiType)
	c.JSON(200, data)
}

func (api *ApiV1) startLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "входной параметр - путь к логфайлу",
	})
}

func (api *ApiV1) stopLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "остановить запись всех логов",
	})
}

func (api *ApiV1) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Информация о подключенном устройстве: статус подключения, IP, время подключения, число переданных пакетов",
	})
}

func (api *ApiV1) config(c *gin.Context) {
	c.JSON(200, gin.H{
		"в разработке": "Конфигурация пути сохранения файла",
	})
}

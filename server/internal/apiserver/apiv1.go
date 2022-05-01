package apiserver

import (
	"csidealer/pkg/databuffer"
	"strconv"
	"github.com/gin-gonic/gin"
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
	p.routGr.GET("/status", p.status)
	p.routGr.GET("/startLog", p.startLog)
	p.routGr.GET("/stopLog", p.stopLog)

	return p
}

func (api *ApiV1) csiLastN(c *gin.Context) {
	csiType := c.Param("type")
	n, _ := strconv.Atoi(c.Query("n"))
	data := api.buf.LastN(n, csiType)
	c.JSON(200, data)
}

func (api *ApiV1) subcarrierLastN(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "амплитудное или фазовое значение - n последних пакетов для конкретной поднесущей",
	})
}

func (api *ApiV1) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Информация о подключенном устройстве: статус подключения, IP, время подключения, число переданных пакетов",
	})
}

func (api *ApiV1) startLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "входной параметр - путь к логфайлу",
	})
}

func (api *ApiV1) stopLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "остановить запись всех логов",
	})
}

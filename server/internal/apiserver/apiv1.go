package apiserver

import (
	"csidealer/pkg/databuffer"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	complex = iota
	abs
	phase
	re
	im
)

type ApiV1 struct {
	routGr *gin.RouterGroup
	buf    *databuffer.PackageBuffer
}

func NewApiV1(routGr *gin.RouterGroup, buf *databuffer.PackageBuffer) *ApiV1 {
	p := new(ApiV1)
	p.routGr = routGr

	p.routGr.GET("/csiLastN", p.csiLastN)
	p.routGr.GET("/subcarrierLastN", p.subcarrierLastN)
	p.routGr.GET("/deviceInfo", p.deviceInfo)
	p.routGr.GET("/startLog", p.startLog)
	p.routGr.GET("/stopLog", p.stopLog)

	return p
}

func (api *ApiV1) csiLastN(c *gin.Context) {
	// Тип может быть complex, abs, phase, re, im
	fmt.Println(api.buf.Data)
	c.JSON(200, gin.H{
		"message": 1,
	})
}

func (api *ApiV1) subcarrierLastN(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "амплитудное или фазовое значение - n последних пакетов для конкретной поднесущей",
	})
}

func (api *ApiV1) deviceInfo(c *gin.Context) {
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

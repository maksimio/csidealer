package apiserver

import "github.com/gin-gonic/gin"

const (
	complex = iota
	abs
	phase
	re
	im
)

func csiLastN(c *gin.Context) {
	// Тип может быть complex, abs, phase, re, im
	c.JSON(200, gin.H{
		"message": "[][]амплитудные или фазовые значения - n последних пакетов",
	})
}

func subcarrierLastN(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "амплитудное или фазовое значение - n последних пакетов для конкретной поднесущей",
	})
}

func deviceInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Информация о подключенном устройстве: статус подключения, IP, время подключения, число переданных пакетов",
	})
}

func startLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "входной параметр - путь к логфайлу",
	})
}

func stopLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "остановить запись всех логов",
	})
}

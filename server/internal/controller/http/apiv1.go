package http

import (
	"csidealer/internal/usecase"
	"strconv"

	// "fmt"
	"github.com/gin-gonic/gin"
)

type ApiV1 struct {
	routGr *gin.RouterGroup
	csiUc  usecase.CsiUC
}

func NewApiV1(rg *gin.RouterGroup, uc usecase.CsiUC) *ApiV1 {
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
	err := a.csiUc.StopLog()
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (a *ApiV1) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"result": gin.H{
			"islogging":     a.csiUc.IsLog(),
			"tcpClientAddr": a.csiUc.GetTcpRemoteAddr(),
			// "isFileConn": false, // TODO: будет добавлено
		},
	})
}

// ---------------------------------- НЕ ГОТОВО:

func (a *ApiV1) csiLastN(c *gin.Context) {
	csi, err := strconv.ParseUint(c.Param("type"), 10, 8)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	data, err := a.csiUc.GetCsi(uint8(csi), n)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"result":  data,
		})
	}
}

func (a *ApiV1) subcarrierLastN(c *gin.Context) {
	// csiType := c.Param("type")
	// h, _ := strconv.Atoi(c.Query("h"))
	// index, _ := strconv.Atoi(c.Query("index"))
	// n, _ := strconv.Atoi(c.Query("n"))
	c.JSON(200, 1)
}

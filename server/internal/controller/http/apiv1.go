package http

import (
	"csidealer/internal/usecase"
	"github.com/gin-gonic/gin"
	"strconv"
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
			// "tcpConnStartTime"
			// "isFileConn": false, // TODO: будет добавлено
			"csiPackageCount": a.csiUc.GetCsiPackageCount(),
			"csiPackageMaxCount": a.csiUc.GetCsiPackageMaxCount(),
			"csiFilter": gin.H{
				"payloadLenMin": 4,
				"payloadLenMax": 4,
				"nr": 4,
				"nc": 4,
				"nTones": 4,
			},
		},
	})
}

func (a *ApiV1) csiLastN(c *gin.Context) {
	csiType, err := strconv.ParseUint(c.Param("type"), 10, 8)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	data, err := a.csiUc.GetCsi(uint8(csiType), n)
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
	csiType, err := strconv.ParseUint(c.Param("type"), 10, 8)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	h, err := strconv.Atoi(c.Query("h"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	i, err := strconv.Atoi(c.Query("i"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}

	data, err := a.csiUc.GetSubcarrier(uint8(csiType), n, h, i)
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

// ---------------------------------- НЕ ГОТОВО:

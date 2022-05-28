package http

import (
	"csidealer/internal/usecase"
	"fmt"
	"strconv"

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
	a.routGr.GET("/csi_last_n/:type", a.csiLastN)
	a.routGr.GET("/subcarrier_last_n/:type", a.subcarrierLastN)
	a.routGr.GET("/status", a.status)

	a.routGr.PUT("/start_log", a.startLog)
	a.routGr.PUT("/stop_log", a.stopLog)
	a.routGr.PUT("/filter_state", a.setFilterState)
	a.routGr.PUT("/filter_limits", a.setFilterLimits)
}

func (a *ApiV1) setFilterLimits(c *gin.Context) {
	_, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()

	intValue, err := strconv.Atoi(c.DefaultQuery("payloadLenMin", strconv.Itoa(int(payloadLenMin))))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	payloadLenMin = uint16(intValue)

	intValue, err = strconv.Atoi(c.DefaultQuery("payloadLenMax", strconv.Itoa(int(payloadLenMax))))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	payloadLenMax = uint16(intValue)

	intValue, err = strconv.Atoi(c.DefaultQuery("nr", strconv.Itoa(int(nr))))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	nr = uint8(intValue)
	intValue, err = strconv.Atoi(c.DefaultQuery("nc", strconv.Itoa(int(nc))))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	nc = uint8(intValue)

	intValue, err = strconv.Atoi(c.DefaultQuery("nTones", strconv.Itoa(int(nTones))))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
		return
	}
	nTones = uint8(intValue)

	state, _, _, _, _, _ := a.csiUc.GetPackageFilterLimits()
	fmt.Println(state, payloadLenMin, payloadLenMax, nr, nc, nTones)
	a.csiUc.SetPackageFilterLimits(state, payloadLenMin, payloadLenMax, nr, nc, nTones)
	c.JSON(200, gin.H{"success": true})
}

func (a *ApiV1) setFilterState(c *gin.Context) {
	state, err := strconv.ParseBool(c.Query("state"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		_, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()
		a.csiUc.SetPackageFilterLimits(state, payloadLenMin, payloadLenMax, nr, nc, nTones)
		c.JSON(200, gin.H{"success": true})
	}
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
	isActive, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()
	c.JSON(200, gin.H{
		"success": true,
		"result": gin.H{
			"serverVersion":      "0.5.1",
			"islogging":          a.csiUc.IsLog(),
			"tcpClientAddr":      a.csiUc.GetTcpRemoteAddr(),
			"csiPackageCount":    a.csiUc.GetCsiPackageCount(),
			"csiPackageMaxCount": a.csiUc.GetCsiPackageMaxCount(),
			"csiFilter": gin.H{
				"isActive":      isActive,
				"payloadLenMin": payloadLenMin,
				"payloadLenMax": payloadLenMax,
				"nr":            nr,
				"nc":            nc,
				"nTones":        nTones,
			},
			// "tcpConnStartTime"
			// "isFileConn": false, // TODO: будет добавлено
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

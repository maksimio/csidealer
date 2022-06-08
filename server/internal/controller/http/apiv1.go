package http

import (
	"csidealer/internal/usecase"
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
	a.routGr.GET("/status", a.status)

	data := a.routGr.Group("/csi")
	data.GET("/last_n/:type", a.csiLastN)
	data.GET("/subcarrier_last_n/:type", a.subcarrierLastN)

	log := a.routGr.Group("/log")
	log.PUT("/start", a.startLog)
	log.PUT("/stop", a.stopLog)

	filter := a.routGr.Group("/filter")
	filter.GET("/state")
	filter.PATCH("/state", a.setFilterState)
	filter.GET("/limits")
	filter.PATCH("/limits", a.setFilterLimits)
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

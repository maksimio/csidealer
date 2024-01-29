package http

import (
	"csidealer/internal/services"

	"github.com/gin-gonic/gin"
)

type ApiV1 struct {
	routGr *gin.RouterGroup
	csiUc  services.CsiUC
}

func NewApiV1(rg *gin.RouterGroup, uc services.CsiUC) *ApiV1 {
	return &ApiV1{
		routGr: rg,
		csiUc:  uc,
	}
}

func (a *ApiV1) Register() {
	// a.routGr.GET("/status", a.status)

	csi := a.routGr.Group("/csi")
	csi.GET("/last_n/:type", a.csiLastN)
	csi.GET("/subcarrier_last_n/:type", a.subcarrierLastN)
	csi.GET("/state") // TODO число пакетов, размер хранилища и т.д.

	log := a.routGr.Group("/log")
	log.GET("/start", a.startLog) // TODO переделать на patch - это касается и других запросов
	log.GET("/stop", a.stopLog)   // TODO переделать на patch
	log.GET("/state", a.stateLog)

	filter := a.routGr.Group("/filter")
	filter.PATCH("/start")                   // TODO
	filter.PATCH("/stop")                    // TODO
	filter.GET("/state")                     // TODO
	filter.PUT("/state", a.setFilterState)   // UPDATE
	filter.PUT("/limits", a.setFilterLimits) // UPDATE

	devices := a.routGr.Group("/devices")
	devices.GET("/tcp_client_ip", a.tcpClientIp)
	devices.GET("/list_info", a.deviceListInfo) // TODO
	devices.POST("/:id")                        // TODO
	devices.DELETE("/:id")                      // TODO
	devices.PATCH("/connect/:id")               // TODO
	devices.PATCH("/disconnect/:id")            // TODO
	devices.PATCH("/send/start/:id")            // TODO
	devices.PATCH("/send/stop/:id")             // TODO
	devices.PATCH("/client/start/:id")          // TODO
	devices.PATCH("/client/stop/:id")           // TODO
}

func (a *ApiV1) tcpClientIp(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "result": a.csiUc.GetTcpRemoteAddr()})
}

// func (a *ApiV1) status(c *gin.Context) {
// 	isActive, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()
// 	c.JSON(200, gin.H{
// 		"success": true,
// 		"result": gin.H{
// 			"serverVersion":      "0.5.1",
// 			"islogging":          a.csiUc.IsLog(),
// 			"tcpClientAddr":      a.csiUc.GetTcpRemoteAddr(),
// 			"csiPackageCount":    a.csiUc.GetCsiPackageCount(),
// 			"csiPackageMaxCount": a.csiUc.GetCsiPackageMaxCount(),
// 			"csiFilter": gin.H{
// 				"isActive":      isActive,
// 				"payloadLenMin": payloadLenMin,
// 				"payloadLenMax": payloadLenMax,
// 				"nr":            nr,
// 				"nc":            nc,
// 				"nTones":        nTones,
// 			},
// 			// "tcpConnStartTime"
// 			// "isFileConn": false, // TODO: будет добавлено
// 		},
// 	})
// }

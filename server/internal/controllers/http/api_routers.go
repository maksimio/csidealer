package http

import "github.com/gin-gonic/gin"

func (h *HttpController) reconnectRouters(c *gin.Context) {
	if err := h.routerConnectorService.Reconnect(); err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (h *HttpController) startCsiTransmit(c *gin.Context) {
	if err := h.routerConnectorService.Start(); err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (h *HttpController) stopCsiTransmit(c *gin.Context) {
	if err := h.routerConnectorService.Stop(); err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (h *HttpController) routersStatus(c *gin.Context) {
	if err := h.routerConnectorService.Stop(); err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

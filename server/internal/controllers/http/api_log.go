package http

import "github.com/gin-gonic/gin"

func (h *HttpController) startLog(c *gin.Context) {
	filepath := c.Query("filepath")
	err := h.rawWriterService.Start(filepath)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (h *HttpController) stopLog(c *gin.Context) {
	err := h.rawWriterService.Stop()
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (h *HttpController) stateLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"result": gin.H{
			"start_ts":         h.rawWriterService.StartTime,
			"is_open":          h.rawWriterService.IsOpen,
			"write_byte_count": h.rawWriterService.WriteByteCount,
			"package_count":    h.rawWriterService.WritePackageCount,
		},
	})
}

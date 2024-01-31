package http

import "github.com/gin-gonic/gin"

func (a *ApiV1) startLog(c *gin.Context) {
	filepath := c.Query("filepath")
	err := a.rawWriterService.Start(filepath)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (a *ApiV1) stopLog(c *gin.Context) {
	err := a.rawWriterService.Stop()
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func (a *ApiV1) stateLog(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"result": gin.H{
			"start_ts":         a.rawWriterService.StartTime,
			"is_open":          a.rawWriterService.IsOpen,
			"write_byte_count": a.rawWriterService.WriteByteCount,
			"package_count":    a.rawWriterService.WritePackageCount,
		},
	})
}

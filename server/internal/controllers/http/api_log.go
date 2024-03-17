package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (h *HttpController) logStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"result": gin.H{
			"start_ts":         h.rawWriterService.StartTime,
			"is_open":          h.rawWriterService.IsWriting,
			"write_byte_count": h.rawWriterService.WriteByteCount,
			"package_count":    h.rawWriterService.WritePackageCount,
		},
	})
}

func (h *HttpController) setMark(c *gin.Context) {
	id := c.Query("id")
	text := c.Query("text")
	isActive, err := strconv.ParseBool(c.Query("is_active"))
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	}

	err = h.rawWriterService.SetMark(id, text, isActive)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

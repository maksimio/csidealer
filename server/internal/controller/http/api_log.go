package http

import "github.com/gin-gonic/gin"

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
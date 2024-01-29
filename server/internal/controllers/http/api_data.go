package http

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

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

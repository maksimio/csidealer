package http

// import (
// 	"github.com/gin-gonic/gin"
// 	"strconv"
// )

// func (a *ApiV1) setFilterLimits(c *gin.Context) {
// 	_, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()

// 	intValue, err := strconv.Atoi(c.DefaultQuery("payloadLenMin", strconv.Itoa(int(payloadLenMin))))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 		return
// 	}
// 	payloadLenMin = uint16(intValue)

// 	intValue, err = strconv.Atoi(c.DefaultQuery("payloadLenMax", strconv.Itoa(int(payloadLenMax))))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 		return
// 	}
// 	payloadLenMax = uint16(intValue)

// 	intValue, err = strconv.Atoi(c.DefaultQuery("nr", strconv.Itoa(int(nr))))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 		return
// 	}
// 	nr = uint8(intValue)
// 	intValue, err = strconv.Atoi(c.DefaultQuery("nc", strconv.Itoa(int(nc))))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 		return
// 	}
// 	nc = uint8(intValue)

// 	intValue, err = strconv.Atoi(c.DefaultQuery("nTones", strconv.Itoa(int(nTones))))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 		return
// 	}
// 	nTones = uint8(intValue)

// 	state, _, _, _, _, _ := a.csiUc.GetPackageFilterLimits()
// 	a.csiUc.SetPackageFilterLimits(state, payloadLenMin, payloadLenMax, nr, nc, nTones)
// 	c.JSON(200, gin.H{"success": true})
// }

// func (a *ApiV1) setFilterState(c *gin.Context) {
// 	state, err := strconv.ParseBool(c.Query("state"))
// 	if err != nil {
// 		c.JSON(500, gin.H{"success": false, "message": err.Error()})
// 	} else {
// 		_, payloadLenMin, payloadLenMax, nr, nc, nTones := a.csiUc.GetPackageFilterLimits()
// 		a.csiUc.SetPackageFilterLimits(state, payloadLenMin, payloadLenMax, nr, nc, nTones)
// 		c.JSON(200, gin.H{"success": true})
// 	}
// }

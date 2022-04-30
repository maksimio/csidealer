package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func RunApiServer(port int, buildDir string) {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(buildDir, true)))
	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/csiLastN", csiLastN)
	v1.GET("/subcarrierLastN", subcarrierLastN)
	v1.GET("/deviceInfo", deviceInfo)
	v1.GET("/startLog", startLog)
	v1.GET("/stopLog", stopLog)

	router.Run(":" + fmt.Sprint(port))
}

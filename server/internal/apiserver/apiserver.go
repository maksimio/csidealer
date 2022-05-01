package apiserver

import (
	"csidealer/pkg/csi"
	"csidealer/pkg/databuffer"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"fmt"
)

func RunApiServer(port int, c <-chan csi.CsiPackage, buildDir string) {
	buf := databuffer.NewPackageBuffer(c)
	go buf.Listen()

	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	NewApiV1(v1, buf)

	router.Use(static.Serve("/", static.LocalFile(buildDir, true)))
	router.Run(":" + fmt.Sprint(port))
}

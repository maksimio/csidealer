package apiserver

import (
	"csidealer/pkg/csicore"
	"csidealer/pkg/databuffer"
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func RunApiServer(port int, c <-chan csicore.CsiPackage, buildDir string) {
	buf := databuffer.NewPackageBuffer(c)
	go buf.Listen()

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(buildDir, true)))
	api := router.Group("/api")
	v1 := api.Group("/v1")

	NewApiV1(v1, buf)
	router.Run(":" + fmt.Sprint(port))
}

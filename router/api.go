package router

import (
	. "github.com/HAL-RO-Developer/shutter/controller"
	"github.com/gin-gonic/gin"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/shutter", Shutter)
	api.POST("/upload", UpLoad)

}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}

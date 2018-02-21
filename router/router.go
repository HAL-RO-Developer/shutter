package router

import (
	"github.com/HAL-RO-Developer/shutter/websocket"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/js", "./public/js")
	r.Static("/image", "./public/image")
	r.Static("/css", "./public/css")
	r.GET("/ws", func(c *gin.Context) {
		websocket.GetHandle()(c.Writer, c.Request)
	})

	r.LoadHTMLGlob("view/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	api := r.Group("/api")
	apiRouter(api)
	auth := api.Group("")
	authApiRouter(auth)
	return r

}

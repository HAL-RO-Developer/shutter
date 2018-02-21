package controller

import (
	"net/http"

	"github.com/HAL-RO-Developer/shutter/websocket"
	"github.com/gin-gonic/gin"
)

func Shutter(c *gin.Context) {
	websocket.Shutter()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

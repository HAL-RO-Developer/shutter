package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, _, _ := c.Request.FormFile("file")
	now := strconv.Itoa(int(time.Now().Unix()))
	dirPath := "./public/image"
	if _, err := os.Stat(dirPath); err != nil {
		os.MkdirAll(dirPath, 0777)
	}
	fileName := now + ".jpg"
	filePath := dirPath + "/" + fileName
	out, _ := os.Create(filePath)

	defer out.Close()
	io.Copy(out, file)

	c.JSON(http.StatusOK, gin.H{
		"path": "/image/" + fileName,
	})
}

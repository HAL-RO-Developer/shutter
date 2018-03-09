package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/HAL-RO-Developer/shutter/service"
	"github.com/gin-gonic/gin"
)

var cnt int

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
	err := service.TweetAPI(service.Account.Get("hal_shutter"), fmt.Sprintf("ようこそ%d人目のお客様です!! \n\n[%s]", cnt, time.Now().String()), filePath)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"path": "/image/" + fileName,
	})
}

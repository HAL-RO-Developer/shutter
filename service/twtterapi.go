package service

import (
	"encoding/base64"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/HAL-RO-Developer/shutter/model"
)

func getTwitterApi(a *model.Account) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(a.CustomerKey)
	anaconda.SetConsumerSecret(a.CustomerSecret)
	api := anaconda.NewTwitterApi(a.AccessToken, a.AccessSecret)
	return api
}

func TweetAPI(a *model.Account, body string, imagepath string) error {
	api := getTwitterApi(a)
	media, _ := api.UploadMedia(fileBase64(imagepath))
	v := url.Values{}
	v.Add("media_ids", media.MediaIDString)

	_, err := api.PostTweet(body, v)
	if err != nil {
		return err
	}
	return nil
}
func fileBase64(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)

	return base64.StdEncoding.EncodeToString(data)
}

package websocket

import (
	"fmt"
	"net/http"

	"github.com/trevex/golem"
)

var room = golem.NewRoom()

func GetHandle() func(http.ResponseWriter, *http.Request) {
	return createRouter().Handler()
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.OnConnect(connectHndle)
	router.OnClose(close)
	router.On("shutter", s)
	return router
}
func Shutter() {
	room.Emit("shutter", struct {
		Msg string `json:"msg"`
	}{
		Msg: "on",
	})
}

func s(conn *golem.Connection) {
	Shutter()
}
func close(conn *golem.Connection) {
	fmt.Println("Leave")
	room.Leave(conn)
}

// connection接続時の処理はここに書く
func connectHndle(conn *golem.Connection, http *http.Request) {
	fmt.Println("Join")
	room.Join(conn)
}

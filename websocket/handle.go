package websocket

import (
	"fmt"
	"net/http"

	"github.com/trevex/golem"
)

var conns []golem.Connection

func GetHandle() func(http.ResponseWriter, *http.Request) {
	return createRouter().Handler()
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.OnConnect(connectHndle)
	router.On("shutter", shutter)
	return router
}
func shutter(conn *golem.Connection) {
	for _, val := range conns {
		if *conn != val {
			val.Emit("", "hoge")
		}
	}
}

// connection接続時の処理はここに書く
func connectHndle(conn *golem.Connection, http *http.Request) {
	fmt.Println(http.Host)
	conns = append(conns, *conn)
}

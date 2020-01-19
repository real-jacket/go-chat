package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 升级 http 为 ws
// 定义一个 upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// 用来检测连接的来源
	// 这将允许我们的 React 服务向这里发出请求
	// 现在我们不需要进行任何的检测
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

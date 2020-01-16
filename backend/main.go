package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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

// 定义一个 reader 来监听网 ws 发出的消息
func reader(conn *websocket.Conn) {
	for {
		// 读消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Print(err)
			return
		}
		// 打印消息
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// 定义 Websocket 服务处理函数
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// 讲连接更新为 WebSocket 连接
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// 一直监听 Websocket 连接上传来的消息
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	// 讲 '/ws' 端点见噢给 'serveWs' 函数处理
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

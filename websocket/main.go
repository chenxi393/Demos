package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// 创建HTTP服务器
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
    //websocketClient()
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// 处理WebSocket连接
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Received message:", string(p))

		// 发送消息
		err = conn.WriteMessage(messageType, []byte("Hello, world!"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func websocketClient() {
	// 连接WebSocket服务器
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
    log.Println("connect success")
	// 发送消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, world!"))
	if err != nil {
		log.Fatal(err)
	}

	// 读取消息
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Received message:", string(p), "type: ", messageType)
}

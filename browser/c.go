// main.go
package main

import (
	"log"
	"net/http"

	//"golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
)

//记录已连接的客户端
var clients = make(map[*websocket.Conn]bool)

//广播通道
var broadcast = make(chan Message)

//升级到http连接到websocket协议
var upgrader = websocket.Upgrader{}

//Message 请求结构
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	//创建静态文件服务
	//fs := http.FileServer(http.Dir("views/"))
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	//设置路由和处理连接方法
	http.HandleFunc("/ws", handleConnections)
	//开始接收和处理请求
	go handleMessages()
	//开始监听8080端口
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//升级http连接到websocket协议
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	//函数返回后关闭此连接
	defer ws.Close()
	//注册新客户端
	clients[ws] = true
	for {
		var req Message
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		//发送接受的请求到消息广播通道
		broadcast <- req
	}
}

func handleMessages() {
	for {
		//从消息广播通道接收消息
		req := <-broadcast
		//发送消息到每个已连接客户端
		for client := range clients {
			err := client.WriteJSON(req)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

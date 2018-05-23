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
var broadcast = make(chan Request)

//升级到http连接到websocket协议
var upgrader = websocket.Upgrader{}

//Request 请求结构
type Request struct {
	Type     string `json:"requesttype"`
	Peer     int    `json:"peer"`
	Name     string `json:"username"`
	OpName   string `json:"operatname"`
	OpAmount int    `json:"operatamount"`
}

/*type Request struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}*/

func main() {
	//创建静态文件服务
	fs := http.FileServer(http.Dir("views/"))
	http.Handle("/", fs)
	//设置路由和处理连接方法
	http.HandleFunc("/ws", handleConnections)
	// Start listening for incoming chat messages
	go handleRequest()
	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	clients[ws] = true
	for {
		var req Request // Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- req
	}
}

func handleRequest() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

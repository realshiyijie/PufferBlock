// main.go
package main

import (
	"fmt"
	"log"
	"myrepo/PufferBlock/server/action"
	"net/http"

	"github.com/gorilla/websocket"
	//"golang.org/x/net/websocket"
)

//记录已连接的客户端
var clients = make(map[*websocket.Conn]bool)

//广播通道
var broadcast = make(chan action.Response)

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

//Websockets ...
func main() {
	//创建静态文件服务
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	//设置路由和处理连接方法
	http.HandleFunc("/ws", handleConnections)
	//开始接收和处理请求
	go handleRequest()
	//开始监听8080端口
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("websockets-ListenAndServe: ", err)
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
		var req Request
		err := ws.ReadJSON(&req)
		fmt.Println(req)
		resp, err := req.doSelect()
		//resp := req
		if err != nil {
			log.Printf("websockets-error: %v", err)
			delete(clients, ws)
			break
		}
		//发送接受的请求到消息广播通道
		broadcast <- resp
	}
}

func handleRequest() {
	for {
		//从消息广播通道接收消息
		resp := <-broadcast
		//发送消息到每个已连接客户端
		for client := range clients {
			err := client.WriteJSON(resp)
			if err != nil {
				log.Printf("websockets-error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func (req *Request) doSelect() (action.Response, error) {

	//缺少类型检查
	if req.Type == "" {
		return action.Response{IfSuccessful: false, ErrInfo: "nil type", Result: ""}, nil
	}

	//初始化账户
	if req.Type == "initUser" {
		return action.InitUser(req.Peer, req.Name)
	}

	//进行交易
	if req.Type == "invoke" {

		//拒接非法交易
		if req.Name != req.OpName {
			return action.Response{IfSuccessful: false, ErrInfo: "denied", Result: ""}, nil
		}
		return action.Invoke(req.Peer, req.Name, req.OpName, req.OpAmount)
	}

	//查询账户信息
	if req.Type == "queryUser" {
		return action.QueryUser(req.Peer, req.OpName)
	}

	//查询所有账户信息
	if req.Type == "queryAll" {
		return action.QueryAll(req.Peer)
	}

	//获取指定账户历史信息
	if req.Type == "getHistory" {
		return action.GetHistory(req.Peer, req.OpName)
	}

	//类型错误抛出
	return action.Response{IfSuccessful: false, ErrInfo: "wrong type", Result: ""}, nil
}

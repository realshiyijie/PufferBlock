//Package server ...
package server

import (
	//	"myrepo/PufferBlock/server/action"
	"myrepo/PufferBlock/server/websockets"
)

//主程序入口
func main() {
	//初始化网络
	//	action.Init()

	//建立连接，接受请求并回复
	websockets.Websockets()
}

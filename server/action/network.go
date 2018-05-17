//Package action 定义了初始化网络和操作账本的方法
package action

import (
	"fmt"
	"os/exec"
)

//生成配置文件
func generate() {
	cmd := "make generate"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(string(out))
}

//启动网络
func networkUp() {
	cmd := "make networkup"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	fmt.Println(string(out))
}

//创建通道
func createChannel() {
	cmd := "sh server.sh createChannel"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//更新锚节点
func updateAnchorPeers() {
	cmd := "sh server.sh updateAnchorPeers"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//安装链码
func installChaincode() {
	cmd := "sh server.sh installChaincode"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//实例化链码
func instantiateChaincode() {
	cmd := "sh server.sh instantiateChaincode"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试初始化用户
func chaincodeInitUser() {
	cmd := "sh server.sh chaincodeInitUser"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试查询
func chaincodeQuery() {
	cmd := "sh server.sh chaincodeQuery"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试交易
func chaincodeInvoke() {
	cmd := "sh server.sh chaincodeInvoke"
	outAsBytes, _ := exec.Command("/bin/sh", "-c", cmd).Output()
	out := string(outAsBytes)
	fmt.Println(string(out))
}

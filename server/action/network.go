//Package action 定义了初始化网络和操作账本的方法
package action

import (
	"fmt"
	"os/exec"
)

var networkScript = "network.sh"

//生成配置文件
func generate() {
	cmd := subCommand + " " + networkScript + " " + "generate"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//启动网络并初始化
func networkUpAndInitByCli() {
	cmd := subCommand + " " + networkScript + " " + "networkUpAndInitByCli"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//启动网络
func networkUp() {
	cmd := subCommand + " " + networkScript + " " + "networkUp"
	out, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	fmt.Println(string(out))
}

//创建通道
func createChannel() {
	cmd := subCommand + " " + networkScript + " " + "createChannel"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//更新锚节点
func updateAnchorPeers() {
	cmd := subCommand + " " + networkScript + " " + "updateAnchorPeers"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//安装链码
func installChaincode() {
	cmd := subCommand + " " + networkScript + " " + "installChaincode"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//实例化链码
func instantiateChaincode() {
	cmd := subCommand + " " + networkScript + " " + "instantiateChaincode"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试初始化用户
func chaincodeInitUser() {
	cmd := subCommand + " " + networkScript + " " + "chaincodeInitUser"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试查询
func chaincodeQuery() {
	cmd := subCommand + " " + networkScript + " " + "chaincodeQuery"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试交易
func chaincodeInvoke() {
	cmd := subCommand + " " + networkScript + " " + "chaincodeInvoke"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}
	out := string(outAsBytes)
	fmt.Println(string(out))
}

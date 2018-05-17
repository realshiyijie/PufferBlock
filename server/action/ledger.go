//Package action 定义了初始化网络和操作账本的方法
package action

import (
	"fmt"
	"os/exec"
	"strconv"
)

var command = "/bin/bash"
var commandArg = "-c"
var subCommand = "bash"
var ledgerScript = "ledger.sh"

//初始化账户
func initCC(name string) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + "initCC" + " " + name
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//进行交易
func invokeCC(name string, opName string, opAmount int) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + "invokeCC" + " " + name + " " + opName + " " + strconv.Itoa(opAmount)
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//查询账户信息
func queryCC(function string, opName string) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + "queryCC" + " " + function + " " + opName
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//Package action 定义了初始化和操作网络的方法
package action

import "os/exec"

//初始化账户
func initCC(name string) (Response, error) {
	cmd := "sh server.sh" + " " + "initCC" + " " + name
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

//进行交易
func invokeCC(function string, opName string, opAmount int) (Response, error) {
	cmd := "sh server.sh" + " " + "invokeCC" + " " + function + " " + opName + " " + string(opAmount)
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

//查询账户信息
func queryCC(opName string) (Response, error) {
	cmd := "sh server.sh" + " " + "queryCC" + " " + opName
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

//Package action ...
package action

import "os/exec"

func initCC(name string) (Response, error) {
	cmd := "sh script.sh" + "initCC" + " " + name
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

func invokeCC(function string, argname string, argamount int) (Response, error) {
	cmd := "sh script.sh" + "invokeCC" + " " + function + " " + argname + " " + string(argamount)
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

func queryCC(argname string) (Response, error) {
	cmd := "sh script.sh" + "queryCC" + " " + argname
	outAsBytes, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	return Response{true, "", out}, nil
}

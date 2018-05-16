//Package websockets ...
package websockets

import "myrepo/PufferBlock/server/action"

func (req *Request) doSelect() (action.Response, error) {

	if req.Mode == "" {
		return action.Response{IfSuccessful: false, ErrInfo: "no mode", Result: ""}, nil
	}
	if req.Mode == "initCC" {
		return action.InitCC(req.Name)
	}
	if req.Mode == "queryCC" {
		return action.QueryCC(req.OpName)
	}
	if req.Mode == "invokeCC" {
		return action.InvokeCC(req.Function, req.OpName, req.OpAmount)
	}
	return action.Response{IfSuccessful: false, ErrInfo: "no mode", Result: ""}, nil
}

//Package action ...
package action

//Response ...
type Response struct {
	IfSuccessful bool   `json:"ifsuccessful"`
	ErrInfo      string `json:"errorinfo"`
	Result       string `json:"result"`
}

//InitCC returns ...
func InitCC(name string) (Response, error) {
	return initCC(name)
}

//InvokeCC returns ...
func InvokeCC(function string, argname string, argamount int) (Response, error) {
	return invokeCC(function, argname, argamount)
}

//QueryCC returns ...
func QueryCC(argname string) (Response, error) {
	return queryCC(argname)
}

//Setup returns ...
func Setup() {
	setup()
}

//Shell returns ...
func Shell() {
	shell()
}

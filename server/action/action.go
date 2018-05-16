//Package action 定义了初始化和操作网络的方法
package action

//Response 回复结构
type Response struct {
	IfSuccessful bool   `json:"ifsuccessful"`
	ErrInfo      string `json:"errorinfo"`
	Result       string `json:"result"`
}

//InitCC 初始化账户接口
func InitCC(name string) (Response, error) {
	return initCC(name)
}

//InvokeCC 进行交易接口
func InvokeCC(function string, opname string, opamount int) (Response, error) {
	return invokeCC(function, opname, opamount)
}

//QueryCC 查询账户信息接口
func QueryCC(opname string) (Response, error) {
	return queryCC(opname)
}

//Init 初始化网络接口
func Init() {

	generate()
	networkUp()
}

//Shell 测试用接口
func Shell() {
	shell()
}

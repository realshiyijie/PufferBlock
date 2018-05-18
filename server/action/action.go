//Package action 定义了初始化网络和操作账本的方法
package action

//Response 回复结构
type Response struct {
	IfSuccessful bool   `json:"ifsuccessful"`
	ErrInfo      string `json:"errorinfo"`
	Result       string `json:"result"`
}

//InitUser 初始化账户接口
func InitUser(peer int, name string) (Response, error) {
	return initUser(peer, name)
}

//Invoke 进行交易接口
func Invoke(peer int, name string, opName string, opAmount int) (Response, error) {
	return invoke(peer, name, opName, opAmount)
}

//QueryUser 查询账户信息接口
func QueryUser(peer int, opName string) (Response, error) {
	return queryUser(peer, opName)
}

//QueryAll 查询所有账户信息接口
func QueryAll(peer int) (Response, error) {
	return queryAll(peer)
}

//Init 初始化网络接口
func Init() {

	generate()
	networkUp()

}

//Package action ...
package action

//InitCC returns ...
func InitCC(name string) bool {
	return initCC(name)
}

//InvokeCC returns ...
func InvokeCC(name string, argname string, argamount int) bool {
	return invokeCC(name, argname, argamount)
}

//QueryCC returns ...
func QueryCC(name string, argname string) bool {
	return queryCC(name, argname)
}

//Setup returns ...
func Setup() {
	setup()
}

//Shell returns ...
func Shell() {
	shell()
}

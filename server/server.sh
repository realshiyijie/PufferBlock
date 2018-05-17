#ledger.go操作网络用脚本
NAME=$1

#测试用
function test() {
	echo "this is: \"${NAME}\""
	docker ps > log.txt
}

test
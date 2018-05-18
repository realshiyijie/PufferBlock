#ledger.go操作网络用脚本

#设置全局变量
PEER=$1
COMMAND=$2
ARG3=$3
ARG4=$4
ARG5=$5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
CHANNEL_NAME="mychannel"
DOCKER_PEER_COMMAND="docker exec cli /go/bin/peer -c"
DOCKER_BASH_COMMAND="docker exec cli /bin/bash -c"
PEER_CHAINCODE_COMMAND="peer chaincode"
OEDERER_ADDRESS="orderer.example.com:7050"
ECHO_COMMAND="echo"

#设置环境变量
setGlobals() {

	PEER=$PEER

	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh ${PEER}"
}

#初始化账户
initCC() {
	
	NAME=$ARG3
	
	$DOCKER_PEER_COMMAND $PEER_CHAINCODE_COMMAND invoke -o $OEDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["createCarbonInfo",'\"${NAME}\"',"RMB","100"]}' 2>&1|grep "status"
}

#进行交易
invokeCC() {
	
	NAME=$ARG3
	OPNAME=$ARG4
	AMOUNT=$ARG5
	
	$DOCKER_PEER_COMMAND $PEER_CHAINCODE_COMMAND invoke -o $OEDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["transfer",'\"${NAME}\"','\"${OPNAME}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"
}

#查询账户信息
queryCC() {
	
	FUNCTION=$ARG3
	OPNAME=$ARG4
	
	if [ "${FUNCTION}" == "queryOne" ]; then
		$DOCKER_PEER_COMMAND $PEER_CHAINCODE_COMMAND query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryByOwner",'\"${OPNAME}\"']}' 2>&1|grep "Query Result"
	elif [ "${FUNCTION}" == "queryAll" ]; then
		$DOCKER_PEER_COMMAND $PEER_CHAINCODE_COMMAND query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryAllCarbonInfo"]}' 2>&1|grep "Query Result"
	else
		$ECHO_COMMAND "ledger.sh-check yr query mode"
	fi
}

#帮助1
arg1Help() {

	$ECHO_COMMAND "ledger.sh-check yr operate mode"
}

#帮助5
arg5Help() {

	$ECHO_COMMAND "ledger.sh-check yr chosen peer"
}

#测试用
test() {
	
	$ECHO_COMMAND "peer${PEER}"
	$ECHO_COMMAND "this is: \"${ARG3}\""
	docker images > log.txt
	bash ./logs/logs.sh
	bash ../blockchain/network/logs.sh
}

#选择执行的操作
if [ ${PEER} -ge 0 ]; then
	setGlobals
	if [ "${COMMAND}" == "initCC" ]; then
		initCC
	elif [ "${COMMAND}" == "invokeCC" ]; then
		invokeCC	
	elif [ "${COMMAND}" == "queryCC" ]; then
		queryCC
	elif [ "${COMMAND}" == "test" ]; then
		test
	else
		arg1Help
	fi
elif [ "${COMMAND}" == "test" ]; then
	test
else
	arg5Help
fi

#ledger.go操作网络用脚本

#设置全局变量
COMMAND=$1
ARG2=$2
ARG3=$3
ARG4=$4
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
CHANNEL_NAME="mychannel"
DOCKER_COMMAND="docker exec cli /go/bin/peer -c"
PEER_COMMAND="peer chaincode"
OEDERER_ADDRESS="orderer.example.com:7050"

#初始化账户
function initCC() {
	NAME=$ARG2
	
	$DOCKER_COMMAND $PEER_COMMAND invoke -o $OEDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["createCarbonInfo",'\"${NAME}\"',"RMB","100"]}' 2>&1|grep "status"
}

#进行交易
function invokeCC() {
	NAME=$ARG2
	OPNAME=$ARG3
	AMOUNT=$ARG4

	$DOCKER_COMMAND $PEER_COMMAND invoke -o $OEDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["transfer",'\"${NAME}\"','\"${OPNAME}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"

}

#查询账户信息
function queryCC() {
	FUNCTION=$ARG2
	OPNAME=$ARG3
	
	if [ "${FUNCTION}" == "queryByOwner" ]; then
	$DOCKER_COMMAND $PEER_COMMAND query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryByOwner",'\"${OPNAME}\"']}' 2>&1|grep "Query Result"
elif [ "${FUNCTION}" == "queryAllCarbonInfo" ]; then
	$DOCKER_COMMAND $PEER_COMMAND query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryAllCarbonInfo"]}' 2>&1|grep "Query Result"
else

}

#测试用
function test() {
	NAME=$ARG2
	ECHO_COMMAND="echo"
	$ECHO_COMMAND "this is: \"${NAME}\""
	docker images > log.txt
}

#选择执行的操作
if [ "${COMMAND}" == "initCC" ]; then
	initCC
elif [ "${COMMAND}" == "invokeCC" ]; then
	invokeCC	
elif [ "${COMMAND}" == "queryCC" ]; then
	queryCC
else
	test
fi
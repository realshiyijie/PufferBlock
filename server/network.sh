#network.go初始化网络用脚本

#设置全局变量
ARG1=$1
ARG2=$2
NETWORK_PATH=../blockchain/network
DOCKER_COMPOSE_CMD=docker-compose
SOURCE_CMD=source
CHANNEL_NAME=mychannel

#生成初始配置
generate() {
    
    $SOURCE_CMD $NETWORK_PATH/generateArtifacts.sh

}

#启动网络
networkUp() {

    $DOCKER_COMPOSE_CMD -f $COMPOSE_FILE up -d 2>&1

}

#帮助
printHelp() {

	echo "network.sh-check yr operate mode"

}

#测试用
test() {
	
	NAME=$ARG2
	ECHO_COMMAND="echo"
	
	$ECHO_COMMAND "this is: \"${NAME}\""
	docker images > log.txt
	bash ./logs/logs.sh
	bash ../blockchain/network/logs.sh

}

if [ "${ARG1}" == "test" ]; then
    test
elif [ "${ARG1}" == "setup" ]; then
    generate
    networkUp
else
    printHelp
fi
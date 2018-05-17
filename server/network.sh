#network.go初始化网络用脚本

#设置全局变量
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

generate
networkUp
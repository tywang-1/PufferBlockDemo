#network.go操作网络用脚本

#设置全局变量
OPERATE=$1
ARG2=$2
NETWORK_PATH="../blockchain/network"
DOCKER_COMPOSE_CMD=docker-compose
SOURCE_CMD="source"
CHANNEL_NAME=mychannel
DOCKER_COMPOSE_CLI=../blockchain/network/docker-compose-cli.yaml
DOCKER_COMPOSE=../blockchain/network/docker-compose.yaml
DOCKER_BASH_COMMAND="docker exec cli /bin/bash -c"

#生成初始配置
generate() {
    
    $SOURCE_CMD ${NETWORK_PATH}/generateArtifacts.sh
}

#启动网络并初始化
networkUpAndInitByCli() {

    $DOCKER_COMPOSE_CMD -f $DOCKER_COMPOSE_CLI up -d 2>&1
}

#启动网络
networkUp() {

	$DOCKER_COMPOSE_CMD -f $DOCKER_COMPOSE up -d 2>&1
}

#创建通道
createChannel() {

	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh createChannel"
}

#更新锚节点
updateAnchorPeers() {

	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh updateAnchorPeers"
}

#加入通道
joinChannel() {

	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh joinChannel"
}

#安装链码
installChaincode() {
	
	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh installChaincode"
}

#实例化链码
instantiateChaincode() {
	
	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh installChaincode"
}

#测试链码
chaincodeTest() {
	
	$DOCKER_BASH_COMMAND "bash ./scripts/setGlobals.sh chaincodeTest"
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

#选择执行的操作
if [ "${OPERATE}" == "generate" ]; then
    generate
elif [ "${OPERATE}" == "networkUpAndInitByCli" ]; then
    networkUpAndInitByCli
elif [ "${OPERATE}" == "networkUp" ]; then
	networkUp
elif [ "${OPERATE}" == "createChannel" ]; then
	createChannel
elif [ "${OPERATE}" == "joinChannel" ]; then
	joinChannel
elif [ "${OPERATE}" == "updateAnchorPeers" ]; then
	updateAnchorPeers
elif [ "${OPERATE}" == "installChaincode" ]; then
	installChaincode
elif [ "${OPERATE}" == "instantiateChaincode" ]; then
	instantiateChaincode    
elif [ "${OPERATE}" == "chaincodeTest" ]; then
	chaincodeTest
else 
	printHelp
fi
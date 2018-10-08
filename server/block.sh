PEER=$1
COMMAND=$2
ARG3=$3
ARG4=$4
ARG5=$5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
CHANNEL_NAME="mychannel"
DOCKER_PEER_COMMAND="docker exec cli /go/bin/peer -c"
DOCKER_BASH_COMMAND="docker exec cli /bin/bash -c"
PEER_CHAINCODE_COMMAND="peer channel"
OEDERER_ADDRESS="orderer.example.com:7050"
ECHO_COMMAND="echo"
NETWORK_PATH="../blockchain/network"

#设置环境变量
getBlockInfo() {

	NUMBER=$ARG3

	$DOCKER_PEER_COMMAND $PEER_CHAINCODE_COMMAND fetch $NUMBER height.block -c $CHANNEL_NAME -o $OEDERER_ADDRESS --tls  --cafile $ORDERER_CA

}

getBlockInfo
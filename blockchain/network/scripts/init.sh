#network.sh初始化网络用

#设置全局变量
OPERATE=$1
CHANNEL_NAME="mychannel"
TIMEOUT:="60"
COUNTER=1
MAX_RETRY=5
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
CHAINCODE_FILE=myrepo/PufferBlock/blockchain/chaincode/go/carbonCC
ORDERER_ADDRESS=orderer.example.com:7050

#结果验证
verifyResult () {
	if [ $1 -ne 0 ] ; then
		echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
                echo "================== ERROR !!! FAILED to execute Initial Scenario =================="
		echo
   		exit 1
	fi
}

#设置环境变量
setGlobals () {

    PEER=$1

	if [ $PEER -eq 0 -o $PEER -eq 1 ] ; then
		CORE_PEER_LOCALMSPID="Org1MSP"
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
		if [ $PEER -eq 0 ]; then
			CORE_PEER_ADDRESS=peer0.org1.example.com:7051
		else
			CORE_PEER_ADDRESS=peer1.org1.example.com:7051
		fi
	else
		CORE_PEER_LOCALMSPID="Org2MSP"
		CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
		CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
		if [ $PEER -eq 2 ]; then
			CORE_PEER_ADDRESS=peer0.org2.example.com:7051
		else
			CORE_PEER_ADDRESS=peer1.org2.example.com:7051
		fi
	fi

	env |grep CORE
}

#创建通道
createChannel() {
	setGlobals 0

	peer channel create -o $ORDERER_ADDRESS -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls TRUE --cafile $ORDERER_CA >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Channel creation failed"
	echo "===================== Channel \"$CHANNEL_NAME\" is created successfully ===================== "
	echo
}

#更新锚节点
updateAnchorPeers() {
    PEER=$1
    setGlobals $PEER

    peer channel update -o $ORDERER_ADDRESS -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls TRUE --cafile $ORDERER_CA >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Anchor peer update failed"
	echo "===================== Anchor peers for org \"$CORE_PEER_LOCALMSPID\" on \"$CHANNEL_NAME\" is updated successfully ===================== "
	sleep 5
	echo
}

#选择peer加入通道
joinWithRetry () {
	peer channel join -b $CHANNEL_NAME.block  >&log.txt
	res=$?
	cat log.txt
	if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
		COUNTER=` expr $COUNTER + 1`
		echo "PEER$1 failed to join the channel, Retry after 2 seconds"
		sleep 2
		joinWithRetry $1
	else
		COUNTER=1
	fi
        verifyResult $res "After $MAX_RETRY attempts, PEER$ch has failed to Join the Channel"
}

#加入通道
joinChannel () {
	for ch in 0 1 2 3; do
		setGlobals $ch
		joinWithRetry $ch
		echo "===================== PEER$ch joined on the channel \"$CHANNEL_NAME\" ===================== "
		sleep 2
		echo
	done
}

#安装链码
installChaincode () {
	PEER=$1
	setGlobals $PEER
	peer chaincode install -n mycc -v 1.0 -p $CHAINCODE_FILE >&log.txt
	res=$?
	cat log.txt
        verifyResult $res "Chaincode installation on remote peer PEER$PEER has Failed"
	echo "===================== Chaincode is installed on remote peer PEER$PEER ===================== "
	echo
}

#实例化链码（一次即可）
instantiateChaincode () {
	PEER=$1
	setGlobals $PEER
	peer chaincode instantiate -o $ORDERER_ADDRESS --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR	('Org1MSP.member','Org2MSP.member')" >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Chaincode instantiation on PEER$PEER on channel '$CHANNEL_NAME' failed"
	echo "===================== Chaincode Instantiation on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
	echo
}

#查询账户测试
chaincodeQuery () {
    PEER=$1
    echo "===================== Querying on PEER$PEER on channel '$CHANNEL_NAME'... ===================== "
    setGlobals $PEER
    local rc=1
    local starttime=$(date +%s)
     while test "$(($(date +%s)-starttime))" -lt "$TIMEOUT" -a $rc -ne 0
     do
        sleep 3
        echo "Attempting to Query PEER$PEER ...$(($(date +%s)-starttime)) secs"
        peer chaincode query -C $CHANNEL_NAME -n mycc -c '{"Args":["queryByOwner","test1"]}' >&log.txt
        test $? -eq 0 && VALUE=$(cat log.txt | awk '/Query Result/ {print $NF}')
        test "$VALUE" = "100" && let rc=0
     done
     echo
     cat log.txt
     if test $rc -eq 0 ; then
    	echo "===================== Query on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
     else
    	echo "!!!!!!!!!!!!!!! Query result on PEER$PEER is INVALID !!!!!!!!!!!!!!!!"
          echo "================== ERROR !!! FAILED to execute Initial Scenario =================="
    	echo
    	exit 1
    fi
}

#初始化账户测试
chaincodeInit () {
	PEER=$1
	setGlobals $PEER
	peer chaincode invoke -o $ORDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["createCarbonInfo","test1"]}' >&log.txt
	peer chaincode invoke -o $ORDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["createCarbonInfo","test2"]}' >&log.txt
    res=$?
	cat log.txt
	verifyResult $res "Initial execution on PEER$PEER failed "
	echo "===================== Initial transaction on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
	echo
}

#进行交易测试
chaincodeInvoke () {
	PEER=$1
	setGlobals $PEER
	peer chaincode invoke -o $ORDERER_ADDRESS  --tls TRUE --cafile $ORDERER_CA -C $CHANNEL_NAME -n mycc -c '{"Args":["transfer","test1","test2","10"]}' >&log.txt
	res=$?
	cat log.txt
	verifyResult $res "Invoke execution on PEER$PEER failed "
	echo "===================== Invoke transaction on PEER$PEER on channel '$CHANNEL_NAME' is successful ===================== "
	echo
}

#帮助
printHelp() {
    
    echo "init.sh-check yr operate mode"
}

if [ "${OPERATE}" == "createChannel" ]; then
    ## Create channel
    echo "Creating channel..."
    createChannel
elif [ "${OPERATE}" == "joinChannel" ]; then
    ## Join all the peers to the channel
    echo "Having all peers join the channel..."
    joinChannel
elif [ "${OPERATE}" == "updateAnchorPeers" ]; then
	## Set the anchor peers for each org in the channel
    echo "Updating anchor peers for org1..."
    updateAnchorPeers 0
    echo "Updating anchor peers for org2..."
    updateAnchorPeers 2
elif [ "${OPERATE}" == "installChaincode" ]; then
    ## Install chaincode on Peer0/Org1 and Peer2/Org2
    echo "Installing chaincode on org1/peer0..."
    installChaincode 0
    echo "Install chaincode on org2/peer2..."
    installChaincode 2
    ## Install chaincode on Peer1/Org1 and Peer3/Org2
    echo "Installing chaincode on org1/peer1..."
    installChaincode 1
    echo "Install chaincode on org2/peer3..."
    installChaincode 3
elif [ "${OPERATE}" == "installChaincode" ]; then
    #Instantiate chaincode on Peer2/Org2
    echo "Instantiating chaincode on org2/peer2..."
    instantiateChaincode 2
elif [ "${OPERATE}" == "chaincodeTest" ]; then
    #Init on chaincode on Peer0/Org1
    echo "Initialing chaincode on org1/peer0..."
    chaincodeInit 0
    #Query on chaincode on Peer0/Org1
    echo "Querying chaincode on org1/peer0..."
    chaincodeQuery 0 100
    #Invoke on chaincode on Peer0/Org1
    echo "Sending invoke transaction on org1/peer0..."
    chaincodeInvoke 0
    #Query on chaincode on Peer3/Org2, check if the result is 90
    echo "Querying chaincode on org2/peer3..."
    chaincodeQuery 3 90
else
    printHelp
fi
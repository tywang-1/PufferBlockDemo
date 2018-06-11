#进行交易用脚本
NAME=$1
OPNAME=$2
AMOUNT=$3
#echo '{"Args":["transfer",'\"${NAME}\"','\"${OPNAME}\"','\"${AMOUNT}\"']}' 
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli /go/bin/peer -c peer chaincode invoke -o orderer.example.com:7050  --tls TRUE --cafile $ORDERER_CA -C mychannel -n mycc -c '{"Args":["transfer",'\"${NAME}\"','\"${OPNAME}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"

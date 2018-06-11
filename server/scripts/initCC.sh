#初始化账户用脚本
NAME=$1
ORDERER_CA=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec  cli /go/bin/peer -c peer chaincode invoke -o orderer.example.com:7050  --tls TRUE --cafile $ORDERER_CA -C mychannel -n mycc -c '{"Args":["createCarbonInfo",'\"${NAME}\"',"RMB","100"]}' 2>&1|grep "status" 

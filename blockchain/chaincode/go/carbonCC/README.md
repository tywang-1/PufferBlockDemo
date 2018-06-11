# CarbonCC碳链说明

## 链码说明

```go
//CarbonInfo 账户信息结构
type CarbonInfo struct {
    //账户类型
    Market string `json:"market"`
    //账户额度
    Amount int    `json:"amount"`
}
```

## 链码操作指令示例

*以下示例在`fabric release-v1.0`上测试通过*

### 进入`fabric-cli`容器

#### 进入`fabric-cli`容器并打开命令行

```bash
docker exec -it cli /bin/bash
```

### 指定环境变量

#### 在执行每条`peer chaincode`命令前需要指定*peer*的所在组织*MSPID*、*TLS*根证书、*MSP*配置文件路径及通信地址

1.选择*org1/peer0*执行`peer chaincode`命令

```bash
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_ADDRESS=peer0.org1.example.com:7051
```

2.选择*org1/peer1*执行`peer chaincode`命令

```bash
CORE_PEER_LOCALMSPID="Org1MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
CORE_PEER_ADDRESS=peer1.org1.example.com:7051
```

3.选择*org2/peer2*执行`peer chaincode`命令

```bash
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
CORE_PEER_ADDRESS=peer0.org2.example.com:7051
```

4.选择*org2/peer3*执行`peer chaincode`命令

```bash
CORE_PEER_LOCALMSPID="Org2MSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
CORE_PEER_ADDRESS=peer1.org2.example.com:7051
```

### 执行`peer chaincode`命令

#### 初始化账户

用户向认证中心注册，登记上链之后，为新用户初始化账户信息

```bash
OWNER=FORREST
MARKET=USD
AMOUNT=100
peer chaincode invoke -o orderer.example.com:7050  --tls TRUE --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["createCarbonInfo",'\"${OWNER}\"','\"${MARKET}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"
```

#### 更新账户信息

运行过程中，可以对用户账户信息进行更新

```bash
OWNER=FORREST
MARKET=AUD
AMOUNT=150
peer chaincode invoke -o orderer.example.com:7050  --tls TRUE --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["updateCarbonInfo",'\"${OWNER}\"','\"${MARKET}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"
```

#### 查询账户信息

1.查询所有账户信息

```bash
peer chaincode query -C mychannel -n mycc -c '{"Args":["queryAllCarbonInfo"]}' 2>&1|grep "Query Result"
```

2.查询指定用户账户信息

```bash
OWNER=FORREST
peer chaincode query -C mychannel -n mycc -c '{"Args":["queryByOwner",'\"${OWNER}\"']}' 2>&1|grep "Query Result"
```

3.查询指定类型账户信息(TODO)

```bash

```

4.查询指定额度账户信息(TODO)

```bash

```

#### 进行交易

向指定对象账户进行交易

```bash
OWNER=FORREST
OPPWNER=JENNY
AMOUNT=10
peer chaincode invoke -o orderer.example.com:7050  --tls TRUE --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n mycc -c '{"Args":["transfer",'\"${OWNER}\"','\"${OPOWNER}\"','\"${AMOUNT}\"']}' 2>&1|grep "status"
```
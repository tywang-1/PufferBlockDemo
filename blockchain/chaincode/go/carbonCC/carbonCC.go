//Package main 碳交易链码
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func main() {
	err := shim.Start(new(CarbonCC))
	if err != nil {
		fmt.Printf("carbonCC-Error creating new Smart Contract: %s", err)
	}
}

//CarbonCC 链码结构
type CarbonCC struct {
}

//CarbonInfo 账户信息结构：Market表示账户类型，Amount表示账户额度
type CarbonInfo struct {
	Market string `json:"market"`
	Amount int    `json:"amount"`
}

//Init 链码初始化接口
func (c *CarbonCC) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

//Invoke 链码操作接口
func (c *CarbonCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	function, args := stub.GetFunctionAndParameters()
	if function == "createCarbonInfo" {
		//创建账户
		return c.createCarbonInfo(stub, args)
	} else if function == "queryAllCarbonInfo" {
		//查询全部账户信息
		return c.queryAllCarbonInfo(stub)
	} else if function == "updateCarbonInfo" {
		//更新账户信息
		return c.updateCarbonInfo(stub, args)
	} else if function == "queryByOwner" {
		//查询指定账户信息
		return c.queryByOwner(stub, args)
	} else if function == "queryByMarket" {
		//查询指定类型账户信息
		return c.queryByMarket(stub, args)
	} else if function == "queryByAmount" {
		//查询指定额度账户信息
		return c.queryByAmount(stub, args)
	} else if function == "transfer" {
		//进行交易
		return c.transfer(stub, args)
	} else if function == "getHistoryForOwner" {
		//获取指定账户历史信息
		return c.getHistoryForOwner(stub, args)
	}

	return shim.Error("carbonCC-Invalid Smart Contract function name.")
}

func (c *CarbonCC) createCarbonInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("carbonCC-Incorrect number of arguments. Expecting 3.")
	}

	//owner 账户所有者，1~10位英文字母，不区分大小写
	owner := strings.ToLower(args[0])
	market := args[1]
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	carbonInfo := &CarbonInfo{market, amount}

	carbonInfoCheckAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state.")
	} else if carbonInfoCheckAsBytes != nil {
		return shim.Error("carbonCC-Account already exists.")
	}

	carbonInfoAsBytes, err := json.Marshal(carbonInfo)
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	stub.PutState(owner, carbonInfoAsBytes)

	//rich.comKeys{Market, amount}

	return shim.Success(nil)
}

func (c *CarbonCC) queryAllCarbonInfo(stub shim.ChaincodeStubInterface) peer.Response {

	//按照用户名命名规则选定范围：a~zzzzzzzzzz
	startKey := "a"
	endKey := "zzzzzzzzzz"

	allCarbonInfoIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state by range.")
	}
	defer allCarbonInfoIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	writtenFlag := false
	for allCarbonInfoIterator.HasNext() {
		queryResponse, err := allCarbonInfoIterator.Next()
		if err != nil {
			return shim.Error("carbonCC-" + err.Error())
		}
		if writtenFlag == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Owner\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")
		buffer.WriteString(", \"CarbonInfo\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		writtenFlag = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (c *CarbonCC) updateCarbonInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("carbonCC-Incorrect number of arguments. Expecting 3.")
	}

	owner := strings.ToLower(args[0])
	market := args[1]
	amount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	carbonInfo := &CarbonInfo{market, amount}

	carbonInfoCheckAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state.")
	} else if carbonInfoCheckAsBytes == nil {
		return shim.Error("carbonCC-Account doesn't exist.")
	}

	carbonInfoAsBytes, err := json.Marshal(carbonInfo)
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	stub.PutState(owner, carbonInfoAsBytes)

	return shim.Success(nil)
}

func (c *CarbonCC) queryByOwner(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("")
	}

	owner := strings.ToLower(args[0])

	carbonInfoAsBytes, err := stub.GetState(owner)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state.")
	} else if carbonInfoAsBytes == nil {
		return shim.Error("carbonCC-Account doesn't exist.")
	}

	return shim.Success(carbonInfoAsBytes)
}

func (c *CarbonCC) queryByMarket(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (c *CarbonCC) queryByAmount(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	return shim.Success(nil)
}

func (c *CarbonCC) transfer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 3 {
		return shim.Error("carbonCC-Incorrect number of arguments. Expecting 3.")
	}

	transferor := strings.ToLower(args[0])
	transferee := strings.ToLower(args[1])
	transferAmount, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}

	carbonInfoTransfereeCheckAsBytes, err := stub.GetState(transferee)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state.")
	} else if carbonInfoTransfereeCheckAsBytes == nil {
		return shim.Error("carbonCC-Transferee doesn't exist.")
	}

	carbonInfoTransferorCheckAsBytes, err := stub.GetState(transferor)
	if err != nil {
		return shim.Error("carbonCC-Failed to get state.")
	} else if carbonInfoTransferorCheckAsBytes == nil {
		return shim.Error("carbonCC-Transferor doesn't exist.")
	}

	carbonInfoTransferorCheck := &CarbonInfo{}
	json.Unmarshal(carbonInfoTransferorCheckAsBytes, carbonInfoTransferorCheck)
	if carbonInfoTransferorCheck.Amount < transferAmount {
		return shim.Error("carbonCC-Transferor Shortage.")
	}
	carbonInfoTransferor := &CarbonInfo{carbonInfoTransferorCheck.Market, carbonInfoTransferorCheck.Amount - transferAmount}
	carbonInfoTransferorAsBytes, err := json.Marshal(carbonInfoTransferor)
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	stub.PutState(transferor, carbonInfoTransferorAsBytes)

	carbonInfoTransfereeCheck := &CarbonInfo{}
	json.Unmarshal(carbonInfoTransfereeCheckAsBytes, carbonInfoTransfereeCheck)
	carbonInfoTransferee := &CarbonInfo{carbonInfoTransfereeCheck.Market, carbonInfoTransfereeCheck.Amount + transferAmount}
	carbonInfoTransfereeAsBytes, err := json.Marshal(carbonInfoTransferee)
	if err != nil {
		return shim.Error("carbonCC-" + err.Error())
	}
	stub.PutState(transferee, carbonInfoTransfereeAsBytes)

	return shim.Success(nil)
}

//获取指定账户历史信息
func (c *CarbonCC) getHistoryForOwner(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("carbonCC-Incorrect number of arguments. Expecting 1.")
	}
	owner := args[0]

	historyIterator, err := stub.GetHistoryForKey(owner)
	if err != nil {
		return shim.Error("carbonCC-Failed to get history for " + owner + ".")
	}
	defer historyIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	writtenFlag := false
	for historyIterator.HasNext() {
		historyData, err := historyIterator.Next()
		if err != nil {
			return shim.Error("carbonCC-" + err.Error())
		}
		if writtenFlag == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(historyData.TxId)
		buffer.WriteString("\"")
		buffer.WriteString(", \"Value\":")
		buffer.WriteString(string(historyData.Value))
		buffer.WriteString(", \"TimeStamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(historyData.Timestamp.Seconds, int64(historyData.Timestamp.Nanos)).String())
		buffer.WriteString("\"")
		buffer.WriteString("}")
		writtenFlag = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

//Package action 定义了操作网络和操作账本的方法
package action

import (
	"fmt"
	"os/exec"
	"strconv"
)

var command = "/bin/bash"
var commandArg = "-c"
var subCommand = "bash"
var ledgerScript = "ledger.sh"

//初始化账户
func initUser(peer int, name string) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + strconv.Itoa(peer) + " " + "initUser" + " " + name
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("ledger-" + err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//进行交易
func invoke(peer int, name string, opName string, opAmount int) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + strconv.Itoa(peer) + " " + "invoke" + " " + name + " " + opName + " " + strconv.Itoa(opAmount)
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("ledger-" + err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//查询账户信息
func queryUser(peer int, opName string) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + strconv.Itoa(peer) + " " + "queryUser" + " " + opName
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("ledger-" + err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//查询所有账户信息
func queryAll(peer int) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + strconv.Itoa(peer) + " " + "queryAll"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("ledger-" + err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

//获取指定账户历史信息
func getHistory(peer int, opName string) (Response, error) {
	cmd := subCommand + " " + ledgerScript + " " + strconv.Itoa(peer) + " " + "getHistory" + " " + opName
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("ledger-" + err.Error())
		return Response{false, err.Error(), ""}, nil
	}
	out := string(outAsBytes)
	fmt.Println(out)
	return Response{true, "", out}, nil
}

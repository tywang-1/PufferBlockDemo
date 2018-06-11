//Package action 定义了操作网络和操作账本的方法
package action

import (
	"fmt"
	"os/exec"
	"strconv"
)

var networkScript = "network.sh"

//初始化网络
func initNetwork(operate string) {
	opSet := []string{"generate", "networkUp", "createChannel", "joinChannel", "updateAnchorPeers", "installChaincode", "instantiateChaincode", "chaincodeTest"}

	if operate == "default" {
		for n, op := range opSet {
			fmt.Println("No." + strconv.Itoa(n) + "step:" + " " + op + "...")
			cmd := subCommand + " " + networkScript + " " + op
			outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
			if err != nil {
				fmt.Println("network-" + err.Error())
			}

			out := string(outAsBytes)
			fmt.Println(out)
		}
	} else {
		cmd := subCommand + " " + networkScript + " " + operate
		outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
		if err != nil {
			fmt.Println("network-" + err.Error())
		}

		out := string(outAsBytes)
		fmt.Println(out)
	}
}

//生成配置文件
func generate() {

	cmd := subCommand + " " + networkScript + " " + "generate"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//启动网络并初始化
func networkUpAndInitByCli() {

	cmd := subCommand + " " + networkScript + " " + "networkUpAndInitByCli"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//启动网络
func networkUp() {

	cmd := subCommand + " " + networkScript + " " + "networkUp"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(out)
}

//创建通道
func createChannel() {

	cmd := subCommand + " " + networkScript + " " + "createChannel"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//加入通道
func joinChannel() {

	cmd := subCommand + " " + networkScript + " " + "joinChannel"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//更新锚节点
func updateAnchorPeers() {

	cmd := subCommand + " " + networkScript + " " + "updateAnchorPeers"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//安装链码
func installChaincode() {

	cmd := subCommand + " " + networkScript + " " + "installChaincode"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//实例化链码
func instantiateChaincode() {

	cmd := subCommand + " " + networkScript + " " + "instantiateChaincode"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

//测试链码
func chaincodeTest() {

	cmd := subCommand + " " + networkScript + " " + "chaincodeTest"
	outAsBytes, err := exec.Command(command, commandArg, cmd).Output()
	if err != nil {
		fmt.Println("network-" + err.Error())
	}

	out := string(outAsBytes)
	fmt.Println(string(out))
}

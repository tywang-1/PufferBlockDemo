//Package websockets 定义了网络通信和接收以及发送消息
package websockets

import "myrepo/PufferBlock/server/action"

//分类解析请求
func (req *Request) doSelect() (action.Response, error) {

	//缺少类型检查
	if req.Type == "" {
		return action.Response{IfSuccessful: false, ErrInfo: "nil type", Result: ""}, nil
	}

	//初始化账户
	if req.Type == "initUser" {
		return action.InitUser(req.Peer, req.Name)
	}

	//进行交易
	if req.Type == "invoke" {

		//拒接非法交易
		if req.Name == req.OpName {
			return action.Response{IfSuccessful: false, ErrInfo: "denied", Result: ""}, nil
		}
		return action.Invoke(req.Peer, req.Name, req.OpName, req.OpAmount)
	}

	//查询账户信息
	if req.Type == "queryUser" {
		return action.QueryUser(req.Peer, req.OpName)
	}

	//查询所有账户信息
	if req.Type == "queryAll" {
		return action.QueryAll(req.Peer)
	}

	//获取指定账户历史信息
	if req.Type == "getHistory" {
		return action.GetHistory(req.Peer, req.OpName)
	}

	//类型错误抛出
	return action.Response{IfSuccessful: false, ErrInfo: "wrong type", Result: ""}, nil
}

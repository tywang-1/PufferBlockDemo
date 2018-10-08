//Package main ...
package main

import (
	"myrepo/PufferBlock/server/action"
	"myrepo/PufferBlock/server/webserver"
	"myrepo/PufferBlock/server/websockets"
	"time"
)

//主程序入口
func main() {
	//初始化网络
	//action.Init()
	//action.GetBlockInfo(0,0)
	//action.InitUser(3, "test")
	//action.InitUser(0, "bbb")
	//action.InitUser(0, "ccc")
	//action.InitUser(0, "ddd")

	//time.Sleep(time.Second * 3)
	//action.QueryAll(0)
	//action.GetHistory(0,"a")
	//建立连接，接受请求并回复
	websockets.Websockets()
	//websockets.Websockets1()
    //testAction()
	//测试
	//testAction
	//testWebserver()
}

//测试链码操作
func testAction() {
  
	action.InitUser(0, "a")
	action.InitUser(2, "b")
	//time.Sleep(time.Second * 3)
	//action.QueryAll(3)
	//action.Invoke(3, "a", "b", 10)
	time.Sleep(time.Second * 3)
	//action.QueryAll(1)
    //action.QueryUser(1, "a")
	//action.QueryUser(2, "b")
	//action.QueryUser(3, "b")
}

func testWebserver() {

	webserver.Webserver()
}

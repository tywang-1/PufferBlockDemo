//Package websockets 定义了网络通信和接收以及发送消息
package websockets

import (
	"fmt"
	"log"
	"myrepo/PufferBlock/server/action"
	"net/http"

	"github.com/gorilla/websocket"
	//"golang.org/x/net/websocket"
)

//记录已连接的客户端
var clients = make(map[*websocket.Conn]bool)

//广播通道
var broadcast = make(chan action.Response)

//升级到http连接到websocket协议
var upgrader = websocket.Upgrader{}

//Request 请求结构
type Request struct {
	Type     string `json:"requesttype"`
	Peer     int    `json:"peer"`
	Name     string `json:"username"`
	OpName   string `json:"operatename"`
	OpAmount int    `json:"operateamount"`
}

//Websockets ...
func Websockets() {
	//创建静态文件服务
	fs := http.FileServer(http.Dir("../browser/"))
	http.Handle("/", fs)
	//设置路由和处理连接方法
	http.HandleFunc("/ws", handleConnections)
	//开始接收和处理请求
	go handleRequest()
	//开始监听8080端口
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("websockets-ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//升级http连接到websocket协议
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	//函数返回后关闭此连接
	defer ws.Close()
	//注册新客户端
	clients[ws] = true
	for {
		var req Request
		err := ws.ReadJSON(&req)
		if err != nil {
			log.Printf("websockets-ReadJSON error: %v", err)
			delete(clients, ws)
			break
		}
		fmt.Println(req)

		resp, err := req.doSelect()
		if err != nil {
			log.Printf("websockets-doSelect error: %v", err)
		}
		//resp := req
		fmt.Println(resp)

		//发送接受的请求到消息广播通道
		broadcast <- resp
	}
}

func handleRequest() {
	for {
		//从消息广播通道接收消息
		resp := <-broadcast
		respStru := action.Response{
			IfSuccessful: resp.IfSuccessful,
			ErrInfo:      resp.ErrInfo,
			Result:       resp.Result,
		}

		//发送消息到每个已连接客户端
		for client := range clients {
			err := client.WriteJSON(respStru)
			if err != nil {
				log.Printf("websockets-WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

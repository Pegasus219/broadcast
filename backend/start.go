package main

import (
	"broadcast/backend/common/utils"
	"broadcast/backend/handle"
	"log"
	"net/http"
)

func main() {
	go startHttp(utils.GetHttpPort())
	go startWebsocket(utils.GetWsPort())
	select {}
}

//启动http网关（接收业务端发出的消息）
func startHttp(port string) {
	log.Println("start http web on port:", port)
	http.HandleFunc("/", handle.HttpWeb)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err.Error())
	}
}

//启动websocket（向接入客户端推送消息）
func startWebsocket(port string) {
	log.Println("start websocket on port:", port)
	http.HandleFunc("/ws", handle.Websocket)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err.Error())
	}
}

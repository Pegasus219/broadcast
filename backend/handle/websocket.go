package handle

import (
	"broadcast/backend/common"
	"github.com/gorilla/websocket"
	"github.com/satori/uuid"
	"net/http"
)

//处理接入ws的客户端的响应
func Websocket(w http.ResponseWriter, r *http.Request) {
	//将http协议升级成websocket协议
	ws := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := ws.Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	//获取接入分组
	group := r.URL.Query().Get("group")

	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	uniqueId := uuid.Must(uuid.NewV4()).String()

	//创建客户端websocket对象
	client := common.NewClient(uniqueId, conn)

	//初始化指定分组下的客户端管理器
	manager := common.InitGroupManager(group)

	//把客户端对象注册到客户端管理者
	manager.RegisterNewClient(client)

	//监听并处理消息推送
	go client.HandlePush()

	//监听并处理客户端离线
	go client.HandleDisconnect(manager)
}

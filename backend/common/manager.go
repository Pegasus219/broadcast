package common

import (
	"broadcast/backend/common/utils"
	"broadcast/backend/message"
)

//ws客户端管理器
type ClientManager struct {
	//分组标签
	group string
	//客户端 map 储存并管理所有的长连接client
	clients map[*Client]bool
	//业务方发送来的的message我们用broadcast来接收，并最后分发给所有的client
	broadcast chan message.MsgInterface
	//新创建的长连接client
	register chan *Client
	//新注销的长连接client
	unregister chan *Client
}

func NewClientManager(group string) *ClientManager {
	return &ClientManager{
		group:      group,
		broadcast:  make(chan message.MsgInterface, utils.GetChannelBuffer()),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//启动websocket用户管理监控
func (m *ClientManager) Start() {
	defer PanicRecover()

	for {
		select {
		//如果有新的连接接入
		case client := <-m.register:
			//把客户端的连接设置为true
			m.clients[client] = true

		//如果连接断开了
		case client := <-m.unregister:
			//关闭client的消息发布通道，撤销已注册到房间的client对象，注意先后顺序
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.message)
			}
			//检查是否需要关闭该clientManager
			if m.IsEmpty() {
				CloseManager(m.group)
				return
			}

		//将同分组内的消息转发给注册的client
		case msg := <-m.broadcast:
			//遍历已经连接的客户端，把消息发送给他们
			for client := range m.clients {
				client.message <- msg
			}
		}
	}
}

//把客户端注册管理者
func (m *ClientManager) RegisterNewClient(client *Client) {
	m.register <- client
}

//从管理器注销客户端
func (m *ClientManager) UnregisterClient(client *Client) {
	m.unregister <- client
}

//检查管理器是否已空
func (m *ClientManager) IsEmpty() bool {
	if len(m.clients) == 0 {
		return true
	}
	return false
}

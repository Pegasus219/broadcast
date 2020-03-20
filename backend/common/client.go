package common

import (
	"github.com/gorilla/websocket"
	"log"
)

//websocket接入客户端
type Client struct {
	//websocket接入用户uuid
	id string
	//连接的socket
	socket *websocket.Conn
	//待发送的消息通道
	message chan *string
}

func NewClient(id string, conn *websocket.Conn) *Client {
	return &Client{
		id:      id,
		socket:  conn,
		message: make(chan *string),
	}
}

//客户端断开websocket连接
func (c *Client) HandleDisconnect(manager *ClientManager) {
	defer PanicRecover()
	defer func() {
		_ = c.socket.Close()
	}()

	for {
		//读取消息
		_, _, err := c.socket.ReadMessage()
		//如果有错误信息（websocket连接断开），就注销这个连接然后关闭
		if err != nil {
			manager.UnregisterClient(c)
			return
		}
	}
}

//向客户端推送消息
func (c *Client) HandlePush() {
	defer PanicRecover()
	defer func() {
		_ = c.socket.Close()
	}()

	for {
		select {
		//读取待发消息
		case msg, ok := <-c.message:
			//如果通道已被关闭，说明websocket连接已断开，退出协程
			if !ok {
				return
			}
			//有消息就写入，发送给客户端
			err := c.socket.WriteMessage(websocket.TextMessage, []byte(*msg))
			if err != nil {
				log.Println(err)
			}
		}
	}
}

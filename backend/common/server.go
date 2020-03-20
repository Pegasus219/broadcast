package common

import (
	"sync"
)

//公共管理者
type managerStruct struct {
	sync.RWMutex
	mc map[string]*ClientManager
}

//初始化公共管理者
var publicManager = &managerStruct{
	mc: map[string]*ClientManager{},
}

//初始化某个分组下的客户端管理器
func InitGroupManager(group string) *ClientManager {
	publicManager.Lock()
	defer publicManager.Unlock()
	if _, has := publicManager.mc[group]; !has {
		publicManager.mc[group] = NewClientManager(group)
		go publicManager.mc[group].Start()
	}
	return publicManager.mc[group]
}

//向指定分组的ws客户端管理器转发消息（如果没有客户端接入，则忽略该消息）
func PushMsg(group, msg string) {
	publicManager.RLock()
	defer publicManager.RUnlock()
	if m, has := publicManager.mc[group]; has {
		m.broadcast <- &msg
	}
}

//关闭空置的客户端管理器
func CloseManager(group string) {
	publicManager.Lock()
	defer publicManager.Unlock()
	delete(publicManager.mc, group)
}

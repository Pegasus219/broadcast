package handle

import (
	"broadcast/backend/common"
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Success bool
}

//接收处理业务端发出的消息
func HttpWeb(w http.ResponseWriter, r *http.Request) {
	group := r.FormValue("group")
	msg := r.FormValue("msg")
	//向指定分组的ws客户端管理器转发消息
	common.PushMsg(group, msg)
	//向业务方ack
	rsp, _ := json.Marshal(HttpResponse{true})
	w.Write(rsp)
}

### 基于websocket消息播报demo

#### 使用方法：
    第一步、运行后端服务 go run backend/start.go
    第二步、浏览器打开前端测试页 frontend/index.html
    第三步、用POST请求向localhost:8080发送消息，参数如下：
| param name | description |
| ----- | ---- |
| group | 分组标签，用于隔离不同业务下的客户端（测试时用test） |
| msg | 消息内容文本 |

#### 关于消息体结构的泛用性（demo示例中为string）
    可在backend/message/structs/下增加自定义的消息体结构
    使其实现message.MsgInterface接口，并替换handle/http.go中的传入参数
package message

type MsgInterface interface {
	//写入
	Inject(message string)
	//提取
	Extract() []byte
}

package utils

import (
	"os"
	"strconv"
)

//获取环境变量
func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

//获取httpWeb的设置端口
func GetHttpPort() string {
	return getEnv("HTTP_PORT", "8080")
}

//获取websocket的设置端口
func GetWsPort() string {
	return getEnv("WS_PORT", "8181")
}

//clientManager中的通道缓冲容量
func GetChannelBuffer() int {
	val := getEnv("CH_BUFFER", "0")
	num, err := strconv.Atoi(val)
	if err != nil || num <= 0 {
		return 50
	}
	return num
}

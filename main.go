package main

import (
	_ "bnwUdp/init"
	"runtime"
	"bnwUdp/server"
)

func main() {
	// 开启多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 启动服务
	server.StartServer()
}

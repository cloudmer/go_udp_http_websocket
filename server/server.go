package server

func StartServer()  {
	// 启动 udp 服务
	go startUdpServer()
	// 启动 http 服务
	go startHttp()

	select {

	}
}

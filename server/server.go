package server

import (
	"fmt"
	"os"
	"bnwUdp/share"
)

func StartServer()  {
	// 启动 udp 服务
	go startUdpServer()
	// 启动 http 服务
	go startHttp()

	fmt.Printf("%c[1;40;32m%s%c[0m\n\n", 0x1B, "北京帮你玩科技有限公司 UDP 设备换绑记录 服务启动成功", 0x1B)
	fmt.Println("进程 ID:",  os.Getpid())
	fmt.Println("UDP Port:",  share.ShareConfig.UdpPort)
	if share.ShareConfig.HttpPort != "" {
		fmt.Println("HTTP Port:",  share.ShareConfig.HttpPort)
		fmt.Println("查看UDP处理情况, 请访问 IP:"+share.ShareConfig.HttpPort)
	}

	select {

	}
}

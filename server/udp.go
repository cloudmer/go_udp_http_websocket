package server

import (
	"net"
	"bnwUdp/share"
	"bnwUdp/library/logger"
	"fmt"
)

// 启动 udp 服务
func startUdpServer()  {
	udpSocket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: share.ShareConfig.UdpPort,
	})
	if err != nil {
		logger.Warning(err.Error())
	}
	defer udpSocket.Close()
	share.ShareUdpServer = udpSocket
	updHandle()
}

// 监听消息
func updHandle()  {
	for {
		data := make([]byte, 1024)
		read, remoteAddr, err := share.ShareUdpServer.ReadFromUDP(data)
		if err != nil {
			fmt.Println("read data ", err)
			continue
		}

		fmt.Println(read, remoteAddr)
		fmt.Printf("%s\n", data)
		send_data := []byte("hi client!")
		_, err = share.ShareUdpServer.WriteToUDP(send_data, remoteAddr)
		if err != nil {
			return
			fmt.Println("send fail!", err)
		}
	}
}
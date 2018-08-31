package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("udpClient")
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 10086,
	})
	if err != nil {
		fmt.Println("connect fail !", err)
		return
	}
	defer socket.Close()

	senddata := []byte("hi server!")
	_, err = socket.Write(senddata)
	if err != nil {
		fmt.Println("send fail !", err)
		return
	}

	data := make([]byte, 10)
	read, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read fail !", err)
		return
	}
	fmt.Println(read, remoteAddr)
	fmt.Printf("%s\n", data)
}

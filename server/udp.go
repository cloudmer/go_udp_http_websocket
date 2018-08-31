package server

import (
	"net"
	"bnwUdp/share"
	"bnwUdp/library/logger"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"strconv"
	"bnwUdp/models"
)

type changeNewDevice struct {
	OldDevice string `json:"old_device"` // 旧设备号
	NewDevice string `json:"new_device"` // 新设备号
	Time 	  int    `json:"time"` // 系统时间
	SecretKey string `json:"secret_key"` // 密钥
}

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
		byte_length, remoteAddr, err := share.ShareUdpServer.ReadFromUDP(data)
		// 去除byte 多余的位置 不然解析不了
		data = data[:byte_length]
		// 记录日志
		logger.Info(remoteAddr.IP.To4().String()+": "+string(data))
		if err != nil {
			continue
		}
		// 广播消息
		wsBroadcastMsg(remoteAddr.String()+": "+string(data))

		// 处理消息
		checkSourceMsg(remoteAddr.IP.To4().String(), data)
	}
}

// 检查来源消息
func checkSourceMsg(remoteAddr string, byteContents []byte)  {
	// byte 转 string
	strContents := string(byteContents)
	if strContents == "" {
		// 记录日志
		logger.Info(remoteAddr+": "+strContents+ " 发来的消息为空 Or 大于 1024 byte")
		// 广播消息
		wsBroadcastMsg(remoteAddr+": "+strContents+ " 发来的消息为空 Or 大于 1024 byte")
		return
	}

	// json 转 结构
	var newDevice changeNewDevice
	if err := json.Unmarshal(byteContents, &newDevice); err != nil {
		logger.Info(remoteAddr+": "+strContents+ " json 格式不正确 "+err.Error())
		wsBroadcastMsg(remoteAddr+": "+strContents+ " json 格式不正确 "+err.Error())
		return
	}

	// 检查值
	if newDevice.OldDevice == "" || newDevice.NewDevice == "" || newDevice.Time <= 0 || newDevice.SecretKey == "" {
		logger.Info(remoteAddr+": "+strContents+ " 数据格式不正确")
		wsBroadcastMsg(remoteAddr+": "+strContents+ " 数据格式不正确")
		return
	}

	// 密钥检查
	str := newDevice.OldDevice + newDevice.NewDevice + strconv.Itoa(newDevice.Time) + share.ShareConfig.Key
	has := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", has)
	secretKey := string([]byte(md5Str)[26:])
	if secretKey != newDevice.SecretKey {
		logger.Info(remoteAddr+": "+strContents+ " secret_key 不正确")
		wsBroadcastMsg(remoteAddr+": "+strContents+ " secret_key 不正确")
		return
	}

	// 准备入库
	logger.Info(remoteAddr+": "+strContents+ " 数据准备入库")
	wsBroadcastMsg(remoteAddr+": "+strContents+ " 数据准备入库")

	// 入库操作
	model := new(models.UdpDeviceChange)
	model.OldDevice = newDevice.OldDevice
	model.NewDevice = newDevice.NewDevice
	err := model.Insert()
	// 入库失败
	if err != nil {
		logger.Debug(remoteAddr+": "+strContents+ " 数据入库失败 "+ err.Error())
		wsBroadcastMsg(remoteAddr+": "+strContents+ " 数据入库失败 "+ err.Error())
		return
	}
	// 入库成功
	logger.Debug(remoteAddr+": "+strContents+ " 数据入库成功～")
	wsBroadcastMsg(remoteAddr+": "+strContents+ " 数据入库成功～")
}
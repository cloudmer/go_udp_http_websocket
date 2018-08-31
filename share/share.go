package share

import (
	"bnwUdp/config"
	"os"
	"database/sql"
	"net"
	"sync"
)

// 全局 config 配置对象
var ShareConfig *config.Config

// 日志文件
var ShareLoggerFile *os.File

// 全局 mysql DB 连接据柄
var ShareDb *sql.DB

// 全局Udp服务
var ShareUdpServer *net.UDPConn

// websocket 客户端 列表
var ShareWsSession sync.Map
package server

import (
	"net/http"
	"bnwUdp/library/logger"
	"github.com/gorilla/websocket"
	"bnwUdp/session"
	"fmt"
)

// http 协议 升级为 websocket 协议
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 10,
	WriteBufferSize: 1024 * 10,
	//允许任何客户端连接
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 将 http 协议 转成 websocket 协议
func websocketServer(writer http.ResponseWriter, request *http.Request)  {
	wsConn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		logger.Error("websocket 连接错误: " + err.Error())
		return
	}
	// 有新的客户端连接到 ws 服务端 添加到客户端列表
	session.WsJoin(wsConn)
	fmt.Println(session.WsList())
	// 监听客户端消息
	go websocketOnMessage(wsConn)
}

// 监听 ws 客户端 发送过来的 消息
func websocketOnMessage(wsConn *websocket.Conn)  {
	for  {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			// 客户端断开链接
			// session 删除
			session.WsLogout(wsConn)
			// 关闭连接
			wsConn.Close()
			break
		}
		wsSendMsg(wsConn, string(message))
	}
}

// 给客户端发送消息
func wsSendMsg(wsConn *websocket.Conn, contents string)  {
	if contents == "" {
		return
	}
	if err := wsConn.WriteMessage(websocket.TextMessage, []byte(contents+"\r\n")); err != nil {
		// 客户端断开连接
		// session 删除
		session.WsLogout(wsConn)
		// 关闭连接
		wsConn.Close()
	}
}

// ws 广播消息
func wsBroadcastMsg(contents string)  {
	for _, wsConn := range session.WsList() {
		wsSendMsg(wsConn, contents)
	}
}

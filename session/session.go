package session

import (
	"github.com/gorilla/websocket"
	"bnwUdp/share"
)

// ws 客户端连接
func WsJoin(ws *websocket.Conn)  {
	share.ShareWsSession.LoadOrStore(ws, ws)
}

// ws 断开连接 列表中删除此连接
func WsLogout(ws *websocket.Conn)  {
	share.ShareWsSession.Delete(ws)
}

// ws 列表
func WsList() []*websocket.Conn {
	wsList := make([]*websocket.Conn, 0)
	share.ShareWsSession.Range(func(key, value interface{}) bool {
		wsList = append(wsList, value.(*websocket.Conn))
		return true
	})
	return wsList
}
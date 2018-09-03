package server

import (
	"net/http"
	"bnwUdp/share"
	"html/template"
)

// 启动 http 服务
func startHttp()  {
	// 检查是否配置 http 端口号
	if share.ShareConfig.HttpPort == "" {
		return
	}

	// bnw 路由地址
	http.HandleFunc("/bnw", websocketServer)
	// 静态文件路由地址
	http.HandleFunc("/static/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, share.ShareConfig.StaticPath + request.URL.Path[7:])
	})
	// 根目录 首页
	http.HandleFunc("/", actionIndex)
	http.ListenAndServe("0.0.0.0:"+share.ShareConfig.HttpPort, nil)
}

// 根目录 首页
func actionIndex(writer http.ResponseWriter, request *http.Request)  {
	template.Must(template.ParseFiles(share.ShareConfig.StaticPath+ "/html/index.html")).Execute(writer, nil)
}
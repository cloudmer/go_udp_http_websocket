package server

import (
	"net/http"
	"os"
	"html/template"
	"bnwUdp/share"
	"fmt"
)

// 启动 http 服务
func startHttp()  {
	// 检查是否配置 http 端口号
	fmt.Println(share.ShareConfig.HttpPort)
	if share.ShareConfig.HttpPort == "" {
		return
	}

	// bnw 路由地址
	http.HandleFunc("/bnw", websocketServer)
	// 静态文件路由地址
	http.HandleFunc("/static/", func(writer http.ResponseWriter, request *http.Request) {
		pwd, _ := os.Getwd()
		http.ServeFile(writer, request, pwd +"/"+ request.URL.Path[1:])
	})
	// 根目录 首页
	http.HandleFunc("/", actionIndex)
	http.ListenAndServe("0.0.0.0:"+share.ShareConfig.HttpPort, nil)
}

// 根目录 首页
func actionIndex(writer http.ResponseWriter, request *http.Request)  {
	pwd, _ := os.Getwd()
	template.Must(template.ParseFiles(pwd+ "/static/html/index.html")).Execute(writer, nil)
}
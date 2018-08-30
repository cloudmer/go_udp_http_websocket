package init

// 初始化
func init()  {
	// 读取 config 配置文件
	option()
	// 检查 runtime 文件夹 没有则生成
	runtimeDir()
	// 创建 程序 logger 文件
	loggerFile()
	// 连接 mysql
	connectMysql()
}
package init

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"bnwUdp/share"
	"bnwUdp/library/utils"
	"database/sql"
	"bnwUdp/library/logger"
	_ "github.com/go-sql-driver/mysql"
)

// 读取 config yaml 文件
func option()  {
	configFilePath := flag.String("f", "", "配置文件")
	flag.Parse()

	if *configFilePath == "" {
		fmt.Println(" -f 配置文件")
		os.Exit(0)
	}

	data, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		fmt.Println("读取配置文件 yaml 错误:", err.Error())
		os.Exit(0)
	}

	if err := yaml.Unmarshal(data, &share.ShareConfig); err != nil {
		fmt.Println("解析配置文件 yaml 错误: ", err.Error())
		os.Exit(0)
	}
}

// runtime 文件夹
func runtimeDir()  {
	// 如果没有 配置 runtime 文件夹 就在当前进程目录下创建 runtime 文件夹
	if share.ShareConfig.RuntimeDir == "" {
		pwd, _ := os.Getwd()
		share.ShareConfig.RuntimeDir = pwd + "/runtime"
	}
	_, err := utils.GenerateDir(share.ShareConfig.RuntimeDir)
	if err != nil {
		fmt.Println("runtime 文件夹创建失败, 请手动创建并赋予可写可读权限", err.Error())
		os.Exit(0)
	}
}

// logger 文件
func loggerFile()  {
	// 如果有 logger 文件 则追究 如果没有 logger 则 创建
	_, err := utils.GenerateFile(share.ShareConfig.RuntimeDir +"/"+ share.ShareConfig.LoggerFileName)
	if err != nil {
		fmt.Println("logger 文件创建失败", err.Error())
		os.Exit(0)
	}
}

// 连接 mysql
func connectMysql()  {
	dataSourceName := share.ShareConfig.Mysql.Username+":"+share.ShareConfig.Mysql.Password+
		"@tcp("+share.ShareConfig.Mysql.Host+":"+share.ShareConfig.Mysql.Port+")/"+
		share.ShareConfig.Mysql.Database+"?charset=utf8"
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Warning(err.Error())
	}

	// 测试连接
	err = db.Ping()
	if err != nil {
		logger.Warning(err.Error())
	}

	// 设置 mysql 空闲最大连接数
	db.SetMaxIdleConns(share.ShareConfig.Mysql.MaxIdle)

	share.ShareDb = db
}
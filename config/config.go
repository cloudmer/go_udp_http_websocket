package config

// 全局配置项 结构体
type Config struct {
	// Mysql 配置
	Mysql struct{
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Database  string `yaml:"database"`
		MaxIdle   int    `yaml:"maxIdle"`
	}

	UdpPort 		int `yaml:"udpPort"`
	HttpPort 		string `yaml:"httpPort"`
	StaticPath		string `yaml:"staticPath"`

	// udp json 加密key
	Key 			string `yaml:"key"`

	// 日志文件夹地址
	RuntimeDir 		string `yaml:"runtime_dir"`

	// 日志文件名
	LoggerFileName 	string `yaml:"logger_file_name"`
}

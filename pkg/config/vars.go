package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 命令行标志定义
var (
	port     = pflag.String("port", "5000", "配置启动监听端口")
	workdir  = pflag.String("workdir", "./", "设置工作目录，用于存放数据库文件, 例如: /var/lib/ip2location/")
	db_type  = pflag.String("db_type", "", "IP 数据库类型，默认空即为IPv4，可选 IPv6")
	db_level = pflag.String("db_level", "", "IP 数据库等级，默认为 DB11，可选 DB1 DB3 DB5 DB9 DB11， 数字越大数据库内容越丰富，相应数据库也就越大")
	token    = pflag.String("token", "", "ip2location lite token")
)

// var (

// // TOKEN   = "domlK3EiUu0anGNXtFhzZZyEZNOmH1eloqsEIqjWovzMY9d2WX78tYx37SZopTzj"
// // WORKDIR = "./"

// // DATABASE_IPTYPE: "IPv6" or ""
// // DB_TYPE = ""
// // DATABASE_LEVEL: "DB1" or "DB3" or "DB5" or "DB9" or "DB11"
// // DB_LEVEL = "DB11"
// )

// BindEnvFor 绑定环境变量
func BindEnvFor() {
	// 绑定所有环境变量
	// viper.AutomaticEnv()
	// 绑定环境变量
	viper.BindEnv("port", "PORT")
	viper.BindEnv("workdir", "WORKDIR")
	viper.BindEnv("db_type", "DB_TYPE")
	viper.BindEnv("db_level", "DB_LEVEL")
	viper.BindEnv("token", "TOKEN")
}

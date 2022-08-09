package config

// package config

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var WORKDIR string

// init 初始化从执行文件所在目录查找配置文件并加载
func init() {
	WORKDIR = viper.GetString("workdir")
	// 配置文件名字，不包含后缀。 此外可以手动指定配置文件格式： viper.SetConfigType("yaml")
	viper.SetConfigName("ip2location")
	// 添加配置搜索的第一个路径，设置为与二进制文件同目录
	viper.AddConfigPath(".")
	viper.AddConfigPath(WORKDIR)
	// 判断加载配置文件是否正确
	if err := viper.ReadInConfig(); err != nil {
		// 判断是否是因为找不到文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 如果是因为找不到文件，则忽略该错误
			ezap.Warn(err)
		} else {
			// 如果是因为文件读取出现错误，则报错退出
			ezap.Fatalf("Read config file failed: %v\n", err)
		}
	}
	// 监听文件修改，热加载配置
	viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	ezap.debugf("Config file changed: %s, %s", e.Name, e.Op)
	// })
}

// LoadConfigs 加载配置
func LoadConfigs() {
	// 加载环境变量
	BindEnvFor()
	// 解析传入参数
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

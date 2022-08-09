package ipquery

import (
	"github.com/fimreal/goutils/ezap"
	mfile "github.com/fimreal/goutils/file"
	httpc "github.com/fimreal/goutils/http"
	mzip "github.com/fimreal/goutils/zip"
	"github.com/spf13/viper"
)

var (
	DB_FILENAME string
	DB_TYPE     string
)

// 预下载数据库，并解压
func GetDB() {
	DB_LEVEL := viper.GetString("db_level")
	DB_TYPE := viper.GetString("db_type")
	WORKDIR := viper.GetString("workdir")
	TOKEN := viper.GetString("token")
	DB_CODE := DB_LEVEL + "LITEBIN" + DB_TYPE
	DB_URL := "https://www.ip2location.com/download/?token=" + TOKEN + "&file=" + DB_CODE
	DB_ZIPFILE := WORKDIR + DB_CODE + ".zip"

	ezap.Debug("DB Download Url: ", DB_URL)
	// ezap.Debug("Token: ", TOKEN)

	// 发现旧文件，则跳过下载
	if mfile.PathExists(DB_ZIPFILE) {
		ezap.Warn("发现数据库文件[", DB_ZIPFILE, "]，跳过下载，如需重新下载数据库请手动删除旧文件")
	} else {
		ezap.Infof("开始下载数据库[%s]]，速度取决于您的网络连接速度", DB_ZIPFILE)
		err := httpc.Download(DB_URL, DB_ZIPFILE)
		if err != nil {
			ezap.Fatalf("下载数据库[%s]出错: %s", DB_ZIPFILE, err)
		}
		ezap.Infof("完成数据库[]下载", DB_ZIPFILE)
	}

	if DB_TYPE == "IPv6" {
		DB_FILENAME = WORKDIR + "IP2LOCATION-LITE-" + DB_LEVEL + ".IPV6.BIN"
	} else {
		DB_FILENAME = WORKDIR + "IP2LOCATION-LITE-" + DB_LEVEL + ".BIN"
	}

	if mfile.PathExists(DB_FILENAME) {
		ezap.Warn("发现数据库文件[", DB_FILENAME, "]，跳过解压缩步骤")
	} else {
		err := mzip.Unzip(DB_ZIPFILE, WORKDIR)
		if err != nil {
			ezap.Fatalf("解压数据库[%s]出错: %s", DB_ZIPFILE, err)
		}
		ezap.Infof("完成数据库[%s]解压缩", DB_ZIPFILE)
	}
	ezap.Infof("配置使用数据文件[%s]", DB_FILENAME)
}

package main

import (
	config "github.com/fimreal/go-ip2location/pkg/config"
	serve "github.com/fimreal/go-ip2location/pkg/http"
	"github.com/fimreal/go-ip2location/pkg/ipquery"

	"github.com/spf13/viper"
)

func main() {
	//
	config.LoadConfigs()
	ipquery.GetDB()
	serve.HandleRequests(serve.IpQuery, ":"+viper.GetString("port"), "/ipquery")
}

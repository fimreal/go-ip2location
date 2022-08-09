package ipquery

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ip2location/ip2location-go/v9"
)

func Query(ip string) (string, error) {
	t := ip2location.OpenTools()

	if !t.IsIPv4(ip) || (!t.IsIPv6(ip) && DB_TYPE == "IPv6") {
		return "", fmt.Errorf("传入 ip[%s] 格式错误", ip)
	}

	db, err := ip2location.OpenDB(DB_FILENAME)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer db.Close()

	results, _ := db.Get_all(ip)

	resMap := mapStruct(results)
	resMap["IP"] = ip

	resJSON, err := json.Marshal(resMap)
	return string(resJSON), err
}

// 过滤数据库不包含的字段
func mapStruct(s interface{}) map[string]interface{} {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	m := make(map[string]interface{})
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		value := v.Field(k).Interface()
		if value != "This parameter is unavailable for selected data file. Please upgrade the data file." {
			m[name] = value
		}
	}
	return m
}

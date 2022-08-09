package serve

import (
	"fmt"
	"net/http"

	"github.com/fimreal/go-ip2location/pkg/ipquery"
	"github.com/fimreal/goutils/ezap"
)

func IpQuery(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		ezap.Println("ParseForm() err: ", err)
		return
	}
	ip := r.FormValue("ip")

	res, err := ipquery.Query(ip)
	if err != nil {
		fmt.Fprintf(w, "查询失败: %v", err)
		ezap.Error("查询失败: ", err)
	}
	// // 返回成功信息
	fmt.Fprintln(w, res)
}

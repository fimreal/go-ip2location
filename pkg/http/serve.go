package serve

import (
	"fmt"
	"net/http"

	"github.com/fimreal/goutils/ezap"
)

// HandleRequests 定义创建服务。
// Handler 为自定义处理器；
// Port 为端口，格式如 [:5000]；
// APIPath 为自定义处理器路径，可以用来加密，格式如：[/abc/123]。
func HandleRequests(Handler func(http.ResponseWriter, *http.Request), Port string, APIPath string) {
	http.HandleFunc(APIPath, Handler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") })
	ezap.Info("Running at ", Port)
	ezap.Info("API: status => /health")
	ezap.Info("API: customAPI => ", APIPath)
	ezap.Fatal(http.ListenAndServe(Port, nil))
}

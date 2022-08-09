package http

import (
	"net/http"

	"github.com/fimreal/goutils/ezap"
)

func HealthServe(HealthServePort string) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})

	ezap.Info("Starting health service")
	ezap.Fatal(http.ListenAndServe(":"+HealthServePort, nil))
}

package broker

import (
	"fmt"
	"net/http"
)

func StartServer(cfg *BrokerConfig) error {
	addr := fmt.Sprintf("%s:%d", cfg.Broker.Service.Host, cfg.Broker.Service.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	return http.ListenAndServe(addr, mux)
}

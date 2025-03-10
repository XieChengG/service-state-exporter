package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"main/collector"
	"net/http"
	"time"
)

func main() {
	// 创建一个自定义注册表
	registry := prometheus.NewRegistry()

	// 注册指标采集器
	registry.MustRegister(collector.NewStateCollector())

	// 配置http路由
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
			Timeout:           5 * time.Second,
		},
	))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<html><body>Service State Exporter</body></html>`))
	})

	// 配置http服务器
	server := &http.Server{
		Addr:         ":10019",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// 启动http服务器
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

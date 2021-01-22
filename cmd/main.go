package main

import (
	"fmt"
	"go-exporter-demo/internal/collectors"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr     string
	endpoint *collectors.RedisEndpoint
)

//GetEnv 功能: 获取外部参数变量
func GetEnv(key string, defaultVale string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultVale
}

//初始化变量
func loadEnv() {
	//设置第三方服务接口的监控信息
	scheme := GetEnv("PROMETHEUS_ENDPOINT_SCHEME", "http")
	host := GetEnv("PROMETHEUS_ENDPOINT_HOST", "192.168.0.4")
	port := GetEnv("PROMETHEUS_ENDPOINT_PORT", "8019")
	metricsPath := "ws/v1/cluster/apps"
	metricsEndpoint := scheme + "://" + host + ":" + port + "/" + metricsPath

	endpoint = &collectors.RedisEndpoint{
		MetricsEndpont: metricsEndpoint,
	}

	log.Println("addr: " + addr)
	fmt.Println(endpoint)
}

func main() {
	loadEnv() //载入变量

	c := collectors.NewCollector(endpoint)
	if c == nil {
		// panic(c)
		log.Fatal("collectors error", c)
		return
	}

	registry := prometheus.Register(c) //注册prometheus
	if registry != nil {
		log.Fatal("registry error")
		return
	}
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9113", nil))
}

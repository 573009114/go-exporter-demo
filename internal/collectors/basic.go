package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
)

//RedisEndpoint 指标结构体
type RedisEndpoint struct {
	MetricsEndpont string
}

//指标结构体，二次开发可根据需要酌情添加
type collector struct {
	endpoint  *RedisEndpoint
	up        *prometheus.Desc
	nodeTotal *prometheus.Desc
}

//指标前缀，用于区分显示指标
const metricsNamespace = "test"

//newFuncMetric 指标描述符
func newFuncMetric(metricName string, docString string) *prometheus.Desc {
	return prometheus.NewDesc(prometheus.BuildFQName(metricsNamespace, "", metricName), docString, nil, nil)
}

//NewCollector 创建新指针对象
func NewCollector(endpoint *RedisEndpoint) *collector {
	return &collector{
		endpoint:  endpoint,
		up:        newFuncMetric("up", "Able to contac t redis"),
		nodeTotal: newFuncMetric("nodeTotal", "total nodes"),
	}
}

//传递结构体中的指标描述符到channel，新增需求则增加对象
func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.up
	ch <- c.nodeTotal
}

//抓取最新的数据，传递给channel，根据需求增加变量和值
func (c *collector) Collect(ch chan<- prometheus.Metric) {
	up := 1.0
	nodeTotal := 5
	ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, up)
	ch <- prometheus.MustNewConstMetric(c.nodeTotal, prometheus.GaugeValue, float64(nodeTotal))
}

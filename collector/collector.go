package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

const (
	namespace = "service"
	subsystem = "state"
)

type stateCollector struct {
	descs   []*prometheus.Desc
	stateFn func() (map[string]int, error)
}

type metricSpec struct {
	name string
	help string
}

var stateMetrics = []metricSpec{
	{"redis", "The state of redis service"},
	{"mysql", "The state of mysql service"},
	{"rabbitmq", "The state of rabbitmq service"},
	{"gitlab", "The state of gitlab service"},
	{"harbor", "The state of harbor service"},
	{"jenkins", "The state of jenkins service"},
	{"nacos", "The state of nacos service"},
	{"mongodb", "The state of mongodb service"},
}

func NewStateCollector() *stateCollector {
	descs := make([]*prometheus.Desc, len(stateMetrics))

	// 遍历 stateMetrics，为每个指标创建一个描述符
	for i, spec := range stateMetrics {
		descs[i] = prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, spec.name),
			spec.help,
			nil,
			nil,
		)
	}
	return &stateCollector{descs: descs, stateFn: getServiceState}
}

// 实现 Collector 接口的 Describe 方法
func (s *stateCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, desc := range s.descs {
		ch <- desc
	}
}

// 实现 Collector 接口的 Collect 方法
func (s *stateCollector) Collect(ch chan<- prometheus.Metric) {
	states, err := s.stateFn()
	if err != nil {
		log.Printf("[ERROR] get service state error: %v", err)
		return
	}

	if len(states) != len(s.descs) {
		log.Printf("[WARN] metrics data mismatch (metric:%d data:%d)",
			len(s.descs),
			len(states),
		)
		return
	}

	for i, desc := range s.descs {
		serviceName := stateMetrics[i].name
		for k, v := range states {
			if serviceName == k {
				ch <- prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					float64(v))
			}
		}
	}
}

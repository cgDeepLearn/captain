package metricsserver

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsapi "k8s.io/metrics/pkg/apis/metrics"

	"captain/pkg/simple/client/monitor"

)

// node metrics definition
const (
	metricsNodeCPUUsage           = "node_cpu_usage"
	metricsNodeCPUTotal           = "node_cpu_total"
	metricsNodeCPUUltilisation    = "node_cpu_utilisation"
	metricsNodeMemoryUsageWoCache = "node_memory_usage_wo_cache"
	metricsNodeMemoryTotal        = "node_memory_total"
	metricsNodeMemoryUltilisation = "node_memory_utilisation"
)

var nodeMetricsNameList = []string{metricsNodeCPUUsage, metricsNodeCPUTotal, metricsNodeCPUUltilisation, metricsNodeMemoryUsageWoCache, metricsNodeMemoryTotal, metricsNodeMemoryUltilisation}

// pod metrics definition
const (
	metricsPodCPUUsage    = "pod_cpu_usage"
	metricsPodMemoryUsage = "pod_memory_usage_wo_cache"
)

func metricsAPISupported(discoveredAPIGroups *metav1.APIGroupList) bool {
	for _, discoveredAPIGroup := range discoveredAPIGroups.Groups {
		if discoveredAPIGroup.Name != metricsapi.GroupName {
			continue
		}
		for _, version := range discoveredAPIGroup.Versions {
			if _, found := supportedMetricsAPIs[version.Version]; found {
				return true
			}
		}
	}
	return false
}

func parseErrorResp(metrics []string, err error) []monitor.Metric {
	var res []monitor.Metric
	for _, metric := range metrics {
		parsedResp := monitor.Metric{MetricName: metric}
		parsedResp.Error = err.Error()
	}
	return res
}

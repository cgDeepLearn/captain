package metricsserver

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsapi "k8s.io/metrics/pkg/apis/metrics"

	"captain/pkg/simple/client/monitor"
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

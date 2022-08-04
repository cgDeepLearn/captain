package metricsserver

import (
	"time"
	"errors"

	"k8s.io/klog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsclient "k8s.io/metrics/pkg/client/clientset/versioned"

	"captain/pkg/simple/client/k8s"
	"captain/pkg/simple/client/monitor"
)

var (
	supportedMetricsAPIs = map[string]bool{
		"v1beta1": true,
	}
)

type metricsServer struct {
	metricsAPIAvailable bool
	metricsClient metricsclient.Interface
	k8s kubernetes.Interface
}

func NewMetricsClient(ki kubernetes.Interface, ko *k8s.KubernetesOptions) monitor.Interface {
	config, err := clientcmd.BuildConfigFromFlags("", ko.KubeConfig)
	if err != nil {
		klog.Error(err)
	}

	discoveryClient := ki.Discovery()
	apiGroups, err := discoveryClient.ServerGroups()
	if err != nil {
		klog.Error(err)
		return nil
	}

	metricsAPIAvailable := metricsAPISupported(apiGroups)

	if !metricsAPIAvailable {
		klog.Warningf("Metrics API not available.")
		return nil
	}

	metricsClient, err := metricsclient.NewForConfig(config)
	if err != nil {
		klog.Error(err)
		return nil
	}

	var ms metricsServer
	ms.k8s = ki
	ms.metricsAPIAvailable = metricsAPIAvailable
	ms.metricsClient = metricsClient

	return ms
}

func (m metricsServer) GetMetric(expr string, ts time.Time) monitor.Metric {
	var resp monitor.Metric
	return resp
}

func (m metricsServer) GetMetricOverTime(expr string, start, end time.Time, step time.Duration) monitor.Metric {
	var resp monitor.Metric

	return resp
}


func (m metricsServer) GetNamedMetrics(metrics []string, ts time.Time, o monitor.QueryOption) []monitor.Metric {
	var resp []monitor.Metric

	options := monitor.NewQueryOptions()
	o.Apply(options)

	if !m.metricsAPIAvailable {
		klog.Warningf("Metrics API not available.")
		return parseErrorResp(metrics, errors.New("Metrics API not available."))
	}

	switch options.Level {
	case monitor.LevelNode:
		// reutrn m.GetNodeLevelNamedMetrics(metrics, ts, opts)
		return resp
	case monitor.LevelPod:
		//reutrn m.GetPodLevelNamedMetrics(metrics, ts, opts)
		return resp
	default:
		return resp
	}
}

func (m metricsServer) GetNamedMetricsOverTime(metrics []string, start, end time.Time, step time.Duration, o monitor.QueryOption) []monitor.Metric {
	var resp []monitor.Metric

	return resp
}

func (m metricsServer) GetMetadata(namespace string) []monitor.Metadata {
	var meta []monitor.Metadata

	return meta
}

func (m metricsServer) GetMetricLabelSet(expr string, start, end time.Time) []map[string]string {
	var res []map[string]string

	return res
}
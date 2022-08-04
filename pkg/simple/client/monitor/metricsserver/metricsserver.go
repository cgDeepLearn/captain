package metricsserver

import (
	"time"
	"errors"
	"context"
	
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsapi "k8s.io/metrics/pkg/apis/metrics"
	metricsv1beta1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
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
		return parseErrorResp(metrics, errors.New("metrics API not available"))
	}

	switch options.Level {
	case monitor.LevelNode:
		return m.GetNodeLevelNamedMetrics(metrics, ts, options)
	case monitor.LevelPod:
		//return m.GetPodLevelNamedMetrics(metrics, ts, options)
		return resp
	default:
		return resp
	}
}

func (m metricsServer) GetNodeLevelNamedMetrics(metrics [] string, ts time.Time, opts *monitor.QueryOptions) []monitor.Metric {
	var resp []monitor.Metric
	nodes, err := m.getNodes()
	if err != nil {
		klog.Errorf("get nodes error: %v\n", err)
		return parseErrorResp(metrics, err)
	}

	status := make(map[string]corev1.NodeStatus)
	for n := range nodes {
		status[n] = nodes[n].Status
	}

	metricsResult, err := m.getNodeMetricsFromMetricsAPI()
	if err != nil {
		klog.Errorf("Get edge node metrics error %v\n", err)
		return parseErrorResp(metrics, err)
	}

	metricsMap := make(map[string]bool)
	for _, m := range metrics {
		metricsMap[m] = true
	}

	nodeMetrics := make(map[string]*monitor.MetricData)
	for _, metricName := range nodeMetricsNameList {
		_, ok := metricsMap[metricName]
		if ok {
			nodeMetrics[metricName] = &monitor.MetricData{MetricType: monitor.MetricTypeVector}
		}
	}

	var usage, capacity corev1.ResourceList
	for _, m := range metricsResult.Items {
		_, ok := nodes[m.Name]
		if !ok {
			continue
		}
		m.Usage.DeepCopyInto(&usage)
		status[m.Name].Capacity.DeepCopyInto(&capacity)
		metricValues := make(map[string]*monitor.MetricValue)

		for _, metricName := range nodeMetricsNameList {
			metricValues[metricName] = &monitor.MetricValue{Metadata: make(map[string]string)}
			metricValues[metricName].Metadata["node"] = m.Name
		}

		for _, addr := range status[m.Name].Addresses {
			if addr.Type == corev1.NodeInternalIP {
				for _, metricName := range nodeMetricsNameList {
					metricValues[metricName].Metadata["host_ip"] = addr.Address
				}
				break
			}
		}

		for k, v := range metricsMap {
			switch k {
			case metricsNodeCPUUsage:
				if v {
					metricValues[metricsNodeCPUUsage].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(usage.Cpu().MilliValue()) / 1000}
				}
			case metricsNodeCPUTotal:
				if v {
					metricValues[metricsNodeCPUTotal].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(capacity.Cpu().MilliValue()) / 1000}
				}
			case metricsNodeCPUUltilisation:
				if v {
					metricValues[metricsNodeCPUUltilisation].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(usage.Cpu().MilliValue()) / float64(capacity.Cpu().MilliValue())}
				}
			case metricsNodeMemoryUsageWoCache:
				if v {
					metricValues[metricsNodeMemoryUsageWoCache].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(usage.Memory().Value())}
				}
			case metricsNodeMemoryTotal:
				if v {
					metricValues[metricsNodeMemoryTotal].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(capacity.Memory().Value())}
				}
			case metricsNodeMemoryUltilisation:
				if v {
					metricValues[metricsNodeMemoryUltilisation].Sample = &monitor.Point{float64(m.Timestamp.Unix()), float64(usage.Memory().Value()) / float64(capacity.Memory().Value())}
				}
			}
		}

		for _, metricName := range nodeMetricsNameList {
			_, ok := metricsMap[metricName]
			if ok {
				resp = append(resp, monitor.Metric{MetricName: metricName, MetricData: *nodeMetrics[metricName]})
			}
		}
	}

	for _, metricName := range nodeMetricsNameList {
		_, ok := metricsMap[metricName]
		if ok {
			resp = append(resp, monitor.Metric{MetricName: metricName, MetricData: *nodeMetrics[metricName]})
		}
	}
	return resp
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

func (m metricsServer) getNodes() (map[string]corev1.Node, error) {
	nodes := make(map[string]corev1.Node)
	corev1Client := m.k8s.CoreV1()
	nodeList, err := corev1Client.Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nodes, err
	}
	for _, node := range nodeList.Items {
		nodes[node.Name] = node
	}
	return nodes, nil
}

func (m metricsServer) getNodeMetricsFromMetricsAPI() (*metricsapi.NodeMetricsList, error) {
	var err error
	mc := m.metricsClient.MetricsV1beta1()
	nm := mc.NodeMetricses()
	versionedMetrics, err := nm.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	metrics := &metricsapi.NodeMetricsList{}
	err = metricsv1beta1.Convert_v1beta1_NodeMetricsList_To_metrics_NodeMetricsList(versionedMetrics, metrics, nil)
	if err != nil {
		return nil, err
	}
	return metrics, nil
}
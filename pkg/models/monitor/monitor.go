package monitor

import (
	"time"

	"k8s.io/client-go/kubernetes"

	"captain/pkg/simple/client/monitor"
)


type MonitorInterface interface {
	GetMetric(expr, namespace string, time time.Time) (monitor.Metric, error)
	GetMetricOverTime(expr, namespace string, start, end time.Time, step time.Duration) (monitor.Metric, error)
	GetNamedMetrics(metrics []string, time time.Time, opt monitor.QueryOption) Metrics
	GetNamedMetricsOverTime(metrics []string, start, end time.Time, step time.Duration, opt monitor.QueryOption) Metrics
	GetMetadata(namespace string) Metadata
	GetMetricLabelSet(metric, namespace string, start, end time.Time) MetricLabelSet
}

type monitorOperator struct {
	promethues monitor.Interface
	metrics monitor.Interface
	k8s kubernetes.Interface
}

func NewMonitorOperator( k8s kubernetes.Interface, monitorClient monitor.Interface, metricsClient monitor.Interface) *monitorOperator {
	return &monitorOperator{
		promethues: monitorClient,
		metrics: metricsClient,
		k8s: k8s,
	}
}

func (mo monitorOperator) GetMetric(expr, namespace string, time time.Time) (monitor.Metric, error) {
	return mo.metrics.GetMetric(expr, time), nil
}

func (mo monitorOperator) GetMetricOverTime(expr, namespace string, start, end time.Time, step time.Duration) (monitor.Metric, error) {
	return mo.metrics.GetMetricOverTime(expr, start, end, step), nil
}

func (mo monitorOperator) GetNamedMetrics(metrics []string, time time.Time, opt monitor.QueryOption) Metrics {
	res := mo.metrics.GetNamedMetrics(metrics, time, opt)
	options := &monitor.QueryOptions{}
	opt.Apply(options)

	return Metrics{Results: res}
}

func (mo monitorOperator) GetNamedMetricsOverTime(metrics []string, start, end time.Time, step time.Duration, opt monitor.QueryOption) Metrics {
	res := mo.metrics.GetNamedMetricsOverTime(metrics, start, end, step, opt)

	return Metrics{Results: res}
}


func (mo monitorOperator) GetMetadata(namespace string) Metadata {
	data := mo.metrics.GetMetadata(namespace)
	return Metadata{Data: data}
}

func (mo monitorOperator) GetMetricLabelSet(metric, namespace string, start, end time.Time) MetricLabelSet {
	data := mo.metrics.GetMetricLabelSet(metric, start, end)
	return MetricLabelSet{Data: data}
}
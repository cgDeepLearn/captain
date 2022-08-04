package monitor

import (
	"captain/pkg/simple/client/monitor"
)

type Metrics struct {
	Results []monitor.Metric `json:"results"`
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

type Metadata struct {
	Data [] monitor.Metadata `json:"data"`
}

type MetricLabelSet struct {
	Data []map[string]string `json:"data"`
}
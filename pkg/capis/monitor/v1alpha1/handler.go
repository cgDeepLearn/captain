package v1alpha1

import (
	
	"regexp"

	"github.com/emicklei/go-restful"
	"k8s.io/client-go/kubernetes"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"captain/pkg/api"
	monitorModel "captain/pkg/models/monitor"
	"captain/pkg/simple/client/monitor"
)

type handler struct {
	k8s kubernetes.Interface
	monitorOp monitorModel.MonitorInterface
	runtimeClient runtimeclient.Client
}


func NewHandler(k8s kubernetes.Interface, monitorClient monitor.Interface, metricsClient monitor.Interface, runtimeClient runtimeclient.Client) *handler {
	return &handler {
		k8s: k8s,
		monitorOp: monitorModel.NewMonitorOperator(k8s, monitorClient, metricsClient),
		runtimeClient: runtimeClient,
	}
}

func (h handler) handleNodeMetricsQuery(req *restful.Request, resp *restful.Response) {
	rp := parseRequestParams(req)
	option, err := makeQueryOptions(rp, LevelNode)
	if err != nil {
		api.HandleBadRequest(resp, nil, err)
		return
	}

	h.handleNamedMetricsQuery(resp, option)

}

func (h handler) handleNamedMetricsQuery(resp *restful.Response, qo queryOptions) {
	var res monitorModel.Metrics
	var metrics []string
	for _, metric := range qo.namedMetrics {
		ok, _ := regexp.MatchString(qo.metricFilter, metric)
		if ok {
			metrics = append(metrics, metric)
		}
	}
	if len(metrics) == 0 {
		resp.WriteAsJson(res)
		return
	}

	res = h.monitorOp.GetNamedMetrics(metrics, qo.time, qo.option)
	resp.WriteAsJson(res)
}




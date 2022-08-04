package v1alpha1

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	restfulspec "github.com/emicklei/go-restful-openapi"

	"captain/pkg/server/runtime"
	"captain/pkg/simple/client/monitor"
	"captain/pkg/constants"
	monitorModel "captain/pkg/models/monitor"
)

const (
	GroupName = "monitor.captain.io"
	VersionName = "v1alpha1"
	respOK = "ok"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: VersionName}

func AddToContainer(c *restful.Container, k8sClient kubernetes.Interface, monitorClient monitor.Interface, metricsClient monitor.Interface, runtimeClient runtimeclient.Client) error {
	ws := runtime.NewWebService(GroupVersion)
	h := NewHandler(k8sClient, monitorClient, metricsClient, runtimeClient)
	
	ws.Route(ws.GET("/nodes/{node}").
		To(h.handleNodeMetricsQuery).
		Doc("Get node-level metric data of the specific node.").
		Param(ws.PathParameter("node", "Node name.").DataType("string").Required(true)).
		Param(ws.QueryParameter("metrics_filter", "The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both node CPU usage and disk usage: `node_cpu_usage|node_disk_size_usage`. View available metrics at [kubesphere.io](https://docs.kubesphere.io/advanced-v2.0/zh-CN/api-reference/monitoring-metrics/).").DataType("string").Required(false)).
		Param(ws.QueryParameter("start", "Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. ").DataType("string").Required(false)).
		Param(ws.QueryParameter("end", "End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. ").DataType("string").Required(false)).
		Param(ws.QueryParameter("step", "Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).").DataType("string").DefaultValue("10m").Required(false)).
		Param(ws.QueryParameter("time", "A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.").DataType("string").Required(false)).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.NodeMetricsTag}).
		Writes(monitorModel.Metrics{}).
		Returns(http.StatusOK, respOK, monitorModel.Metrics{})).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/nodes").
		To(h.handleNodeMetricsQuery).
		Doc("Get node-level metrics data").
		Param(ws.QueryParameter("metrics_filter", "The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both node CPU usage and disk usage: `node_cpu_usage|node_disk_size_usage`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("resources_filter", "The node filter consists of a regexp pattern. It specifies which node data to return. For example, the following filter matches both node i-caojnter and i-cmu82ogj: `i-caojnter|i-cmu82ogj`.").DataType("string").Required(false)).
		Param(ws.QueryParameter("start", "Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. ").DataType("string").Required(false)).
		Param(ws.QueryParameter("end", "End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. ").DataType("string").Required(false)).
		Param(ws.QueryParameter("step", "Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).").DataType("string").DefaultValue("10m").Required(false)).
		Param(ws.QueryParameter("time", "A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.").DataType("string").Required(false)).
		Param(ws.QueryParameter("target", "Sort nodes by the specified metric. Not applicable if **start** and **end** are provided.").DataType("string").Required(false)).
		Param(ws.QueryParameter("order", "Sort order. One of asc, desc.").DefaultValue("desc.").DataType("string").Required(false)).
		Param(ws.QueryParameter("page", "The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data.").DataType("integer").Required(false)).
		Param(ws.QueryParameter("size", "Page size, the maximum number of results in a single page. Defaults to 5.").DataType("integer").Required(false).DefaultValue("5")).
		Param(ws.QueryParameter("type", "The query type. This field can be set to 'rank' for node ranking query or '' for others. Defaults to ''.").DataType("string").Required(false).DefaultValue("")).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.NodeMetricsTag}).
		Writes(monitorModel.Metrics{}).
		Returns(http.StatusOK, respOK, monitorModel.Metrics{})).
		Produces(restful.MIME_JSON)

	c.Add(ws)
	return nil
}
package v1alpha1

import (
	"strconv"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/pkg/errors"

	monitorModel "captain/pkg/models/monitor"
	"captain/pkg/simple/client/monitor"
)

const (
	DefaultStep   = 10 * time.Minute
	DefaultFilter = ".*"
	DefaultOrder  = "desc"
	DefaultPage   = 1
	DefaultSize   = 5

	OperationQuery  = "query"
	OperationExport = "export"

	ComponentEtcd      = "etcd"
	ComponentAPIServer = "apiserver"
	ComponentScheduler = "scheduler"

	ErrNoHit             = "'end' or 'time' must be after the namespace creation time."
	ErrParamConflict     = "'time' and the combination of 'start' and 'end' are mutually exclusive."
	ErrInvalidStartEnd   = "'start' must be before 'end'."
	ErrInvalidPage       = "Invalid parameter 'page'."
	ErrInvalidLimit      = "Invalid parameter 'limit'."
	ErrParameterNotfound = "Parmameter [%s] not found"
)

// type MonitorLevel int

// const (
// 	LevelCluster MonitorLevel  = 1 << iota
// 	LevelNode
// 	LevelNamespace
// 	LevelWorkload
// 	LevelService
// 	LevelPod
// 	LevelContainer
// 	LevelPVC
// 	LevelComponent
// 	LevelIngress
// )


type reqParams struct {
	time                      string
	start                     string
	end                       string
	step                      string
	target                    string
	order                     string
	page                      string
	size                      string
	metricFilter              string
	namespacedResourcesFilter string
	resourceFilter            string
	nodeName                  string
	workspaceName             string
	namespaceName             string
	workloadKind              string
	workloadName              string
	podName                   string
	containerName             string
	pvcName                   string
	storageClassName          string
	componentType             string
	expression                string
	metric                    string
	applications              string
	cluster                   string
	ingress                   string
	job                       string
	services                  string
	duration                  string
	pvcFilter                 string
	queryType                 string
}

type queryOptions struct {
	metricFilter string
	namedMetrics []string

	start time.Time
	end time.Time
	time time.Time
	step time.Duration

	target string
	order string
	page int
	size int

	option monitor.QueryOption
}

func parseRequestParams(req *restful.Request) reqParams {
	var rp reqParams
	rp.time = req.QueryParameter("time")
	rp.start = req.QueryParameter("start")
	rp.end = req.QueryParameter("end")
	rp.step = req.QueryParameter("step")
	rp.target = req.QueryParameter("target")
	rp.order = req.QueryParameter("order")
	rp.page = req.QueryParameter("page")
	rp.size = req.QueryParameter("size")
	rp.namespaceName = req.PathParameter("namespace")
	rp.metricFilter = req.QueryParameter("metrics_filter")
	rp.nodeName = req.PathParameter("node")
	rp.podName = req.PathParameter("pod")
	rp.containerName = req.PathParameter("container")
	rp.metric = req.QueryParameter("metric")
	rp.queryType = req.QueryParameter("type")
	return rp
}

func makeQueryOptions(rp reqParams, level monitor.MonitorLevel) (qo queryOptions, err error) {
	if rp.metricFilter == "" {
		qo.metricFilter = DefaultFilter
	}
	switch level {
	case monitor.LevelCluster:
		qo.option = monitor.ClusterOption{}
		qo.namedMetrics = monitorModel.ClusterMetrics
	case monitor.LevelNode:
		qo.namedMetrics = monitorModel.NodeMetrics
		qo.option = monitor.NodeOption{
			ResourceFilter:   rp.resourceFilter,
			NodeName:         rp.nodeName,
			QueryType:        rp.queryType,
		}
	
	case monitor.LevelNamespace:
		qo.namedMetrics = monitorModel.NamespaceMetrics
		qo.option = monitor.NamespaceOption{
			ResourceFilter:   rp.resourceFilter,
			NamespaceName:    rp.namespaceName,
		}

	case monitor.LevelPod:
		qo.namedMetrics = monitorModel.PodMetrics
		qo.option = monitor.PodOption{
			NamespacedResourcesFilter: rp.namespacedResourcesFilter,
			ResourceFilter:            rp.resourceFilter,
			NodeName:                  rp.nodeName,
			NamespaceName:             rp.namespaceName,
			WorkloadKind:              rp.workloadKind,
			WorkloadName:              rp.workloadName,
			PodName:                   rp.podName,
		}
	}
	
	if rp.start != "" && rp.end != "" {
		startInt, err := strconv.ParseInt(rp.start, 10, 64)
		if err != nil {
			return qo, err
		}
		qo.start = time.Unix(startInt, 0)

		endInt, err := strconv.ParseInt(rp.end, 10, 64)
		if err != nil {
			return qo, err
		}
		qo.end = time.Unix(endInt, 0)

		if rp.step == "" {
			qo.step = DefaultStep
		} else {
			qo.step, err = time.ParseDuration(rp.step)
			if err != nil {
				return qo, err
			}
		}

		if qo.start.After(qo.end) {
			return qo, errors.New(ErrInvalidStartEnd)
		}
	} else if rp.start == "" && rp.end == "" {
		if rp.time == "" {
			qo.time = time.Now()
		} else {
			timeInt, err := strconv.ParseInt(rp.time, 10, 64)
			if err != nil {
				return qo, err
			}
			qo.time = time.Unix(timeInt, 0)
		}
	} else {
		return qo, errors.Errorf(ErrParamConflict)
	}

	if rp.target != "" {
		qo.target = rp.target
		qo.page = DefaultPage
		qo.size = DefaultSize
		qo.order = rp.order
		if rp.order != "asc" {
			qo.order = DefaultOrder
		}
		if rp.page != "" {
			qo.page, err = strconv.Atoi(rp.page)
			if err != nil || qo.page <= 0 {
				return qo, errors.New(ErrInvalidPage)
			}
		}
		if rp.size != "" {
			qo.size, err = strconv.Atoi(rp.size)
			if err != nil || qo.size <= 0 {
				return qo, errors.New(ErrInvalidLimit)
			}
		}
	}

	return qo, nil
}
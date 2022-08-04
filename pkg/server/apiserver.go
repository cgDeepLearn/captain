package server

import (
	"sync"
	"context"
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/endpoints/handlers/responsewriters"
	"k8s.io/klog"
	"net/http"
    urlruntime "k8s.io/apimachinery/pkg/util/runtime"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	captainserverconfig "captain/pkg/server/config"
	monitorv1alpha1 "captain/pkg/capis/monitor/v1alpha1"
	"captain/pkg/simple/client/k8s"
	"captain/pkg/capis/version"
	"captain/pkg/server/filters"
	"captain/pkg/server/request"
	"captain/pkg/simple/client/monitor"
	"captain/pkg/utils/metrics"
)

type APIServer struct {
	ServerCount int

	Server *http.Server

	Config *captainserverconfig.Config

	// webservice container, where all webservice defines
	container *restful.Container

	KubernetesClient k8s.Client

	// monitor
	MonitorClient monitor.Interface
	// metricsserver
	MetricsClient monitor.Interface

	// controller-runtime client
	RuntimeClient runtimeclient.Client
}

var initMetrics sync.Once

type errorResponder struct{}

func (e *errorResponder) Error(w http.ResponseWriter, req *http.Request, err error) {
	klog.Error(err)
	responsewriters.InternalError(w, req, err)
}

func (s *APIServer) PrepareRun(stopCh <-chan struct{}) error {
	s.container = restful.NewContainer()
	//s.container.Filter(logRequestAndResponse)
	// 设定路由为CurlyRouter(快速路由)
	s.container.Router(restful.CurlyRouter{})
	//s.container.RecoverHandler(func(panicReason interface{}, httpWriter http.ResponseWriter) {
	//	logStackOnRecover(panicReason, httpWriter)
	//})

	//s.installCRDAPIs()
	s.installMetricsAPI()
	//s.container.Filter(monitorRequest)

	for _, ws := range s.container.RegisteredWebServices() {
		klog.V(2).Infof("%s", ws.RootPath())
	}

	// container 作为http server 的handler
	s.Server.Handler = s.container

	// 注册服务
	s.installCaptainAPIs()
	// handle chain
	s.buildHandlerChain(stopCh)

	return nil
}

// Install all captain api groups
// Installation happens before all informers start to cache objects, so
//   any attempt to list objects using listers will get empty results.
func (s *APIServer) installCaptainAPIs() {
	urlruntime.Must(version.AddToContainer(s.container, s.KubernetesClient.Discovery()))
	urlruntime.Must(monitorv1alpha1.AddToContainer(s.container, s.KubernetesClient.Kubernetes(), s.MonitorClient, s.MetricsClient, s.RuntimeClient))
}

func (s *APIServer) installMetricsAPI() {
	initMetrics.Do(registerMetrics)
	metrics.Defaults.Install(s.container)
}

//通过WithRequestInfo解析API请求的信息，WithKubeAPIServer根据API请求信息判断是否代理请求给Kubernetes
func (s *APIServer) buildHandlerChain(stopCh <-chan struct{}) {
	requestInfoResolver := &request.RequestInfoFactory{
		APIPrefixes:          sets.NewString("api", "apis", "capis", "capi"),
		GrouplessAPIPrefixes: sets.NewString("api", "capi"),
	}

	handler := s.Server.Handler
	handler = filters.WithKubeAPIServer(handler, s.KubernetesClient.Config(), &errorResponder{})

	handler = filters.WithRequestInfo(handler, requestInfoResolver)

	s.Server.Handler = handler
}

func (s *APIServer) Run(ctx context.Context) (err error) {

	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-ctx.Done()
		_ = s.Server.Shutdown(shutdownCtx)
	}()

	klog.V(0).Infof("Start listening on %s", s.Server.Addr)
	if s.Server.TLSConfig != nil {
		err = s.Server.ListenAndServeTLS("", "")
	} else {
		err = s.Server.ListenAndServe()
	}

	return err
}


package v1alpha1

import (
	"github.com/emicklei/go-restful"
	"k8s.io/klog"
)



type handler struct {}

type HelloResponse struct {
	Message string `json:"message"`
}

func newHandler() handler{
	return handler{}
}

func (h handler) Hello(request *restful.Request, response *restful.Response) {
	username := request.PathParameter("user")
	klog.V(0).Infof("receive hello request %s:", username)
	response.WriteAsJson(HelloResponse{
		Message: "Hello " + username,
	})

}
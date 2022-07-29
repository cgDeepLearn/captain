package v1alpha1

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"captain/pkg/api"
	"captain/pkg/server/runtime"

)

const (
	GroupName = "hello.captain.io"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func AddToContainer(container *restful.Container) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler()
	webservice.Route(webservice.GET("/hello/{user}").
		To(handler.Hello).
		Doc("Say Hello to user").
		Param(webservice.PathParameter("user", "username")).
		Returns(http.StatusOK, api.StatusOK, HelloResponse{}))

	container.Add(webservice)
	return nil
}
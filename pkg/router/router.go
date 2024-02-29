package router

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

type Router struct {
	*http.ServeMux
}

var _ http.Handler = &Router{}

func NewRouter(routes map[string]http.Handler) Router {
	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.Handle(route, otelhttp.WithRouteTag(route, handler))
	}
	return Router{ServeMux: mux}
}

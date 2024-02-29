package server

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"htmx/pkg/router"
	"net/http"
	"time"
)

type Config struct {
	Addr   string
	Routes map[string]http.Handler
}

type Server struct {
	*http.Server
}

func NewServer(config Config) (*Server, error) {
	srv := &Server{
		Server: &http.Server{
			Addr:         config.Addr,
			ReadTimeout:  time.Second,
			WriteTimeout: 10 * time.Second,
			Handler:      otelhttp.NewHandler(router.NewRouter(config.Routes), "/"),
		},
	}
	return srv, nil
}

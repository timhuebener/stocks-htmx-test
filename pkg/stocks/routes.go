package stocks

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"htmx/pkg/stocks/internal/handler"
	"net/http"
)

func (app *Stocks) routes() map[string]http.Handler {
	return map[string]http.Handler{
		"/": otelhttp.NewHandler(handler.IndexHandler{
			Title:    "Stocks",
			BasePath: app.config.BasePath,
		}, "home"),
		"/hello": otelhttp.NewHandler(handler.HelloHandler{}, "hello"),
	}
}

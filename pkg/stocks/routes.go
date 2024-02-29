package stocks

import (
	"htmx/pkg/stocks/internal/handler"
	"net/http"
)

func (app *Stocks) routes() map[string]http.Handler {
	return map[string]http.Handler{
		"/": handler.IndexHandler{
			Title:    "Stocks",
			BasePath: app.config.BasePath,
		},
		"/hello": handler.HelloHandler{},
	}
}

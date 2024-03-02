package http

import (
	"htmx/pkg/otel"
	"net/http"
)

type Handler struct {
	name    string
	handler http.Handler
	tracer  otel.Tracer
}

func New(name string, handler http.Handler) *Handler {
	return &Handler{
		name:    name,
		handler: handler,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), h.name)
	defer span.End()
	r.WithContext(ctx)
	h.handler.ServeHTTP(w, r)
}

var _ http.Handler = &Handler{}

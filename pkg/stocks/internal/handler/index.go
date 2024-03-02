package handler

import (
	"context"
	"html/template"
	"htmx/pkg/otel"
	"htmx/pkg/otel/log"
	"htmx/pkg/stocks/internal/pgdb"
	"net/http"
	"path"
)

type IndexHandler struct {
	Title    string
	BasePath string
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	content := "home.html"
	layout := "index.html"

	tmpl, err := template.ParseGlob(path.Join(h.BasePath, "components", "*.html"))
	tmpl.ParseFiles(path.Join(h.BasePath, "layouts", layout), path.Join(h.BasePath, "pages", content))
	if err != nil {
		log.Error(context.TODO(), "unable to get html template", otel.ErrorMsg.String(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	res, err := pgdb.GetAccounts(ctx)
	if err != nil {
		log.Error(context.TODO(), "not able to get accounts", otel.ErrorMsg.String(err.Error()))
		w.Write([]byte("Something went wrong"))
		return
	}

	data := struct {
		Title    string
		Accounts []pgdb.Account
	}{
		Title:    h.Title,
		Accounts: res,
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Error(context.TODO(), "unable to execute template", otel.ErrorMsg.String(err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var _ http.Handler = &IndexHandler{}

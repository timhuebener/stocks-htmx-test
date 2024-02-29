package handler

import (
	"html/template"
	"net/http"
	"path"
)

type IndexHandler struct {
	Title    string
	BasePath string
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file := "index.html"
	tmpl, err := template.New(file).ParseFiles(path.Join(h.BasePath, file))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	data := struct {
		Title string
	}{
		Title: h.Title,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var _ http.Handler = &IndexHandler{}

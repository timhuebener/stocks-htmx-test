package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../web/templates/index.html")
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching accounts"))
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error fetching accounts"))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

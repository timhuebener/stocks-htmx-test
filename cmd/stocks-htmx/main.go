package main

import (
	"html/template"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	print("Request received\n")
	tmpl, err := template.ParseFiles("../../web/templates/index.html")
	if err != nil {
		print("Error parsing template\n")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		print("Error executing template\n")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", root)
	print("Starting server on port 8080\n")
	http.ListenAndServe(":8080", nil)
}

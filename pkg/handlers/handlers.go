package handlers

import (
	"html/template"
	"htmx/pkg/ledger"
	ledgermodels "htmx/pkg/ledger/models"
	"io"
	"log"
	"net/http"
	"os"
)

func Root(w http.ResponseWriter, r *http.Request) {
	accounts, err := ledger.Accounts("./ledger2023.ledger")
	if err != nil {
		log.Fatalf("Error fetching accounts: %v \n", err)
	}

	tmpl, err := template.ParseFiles("../../web/templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type data struct {
		Accounts []ledgermodels.Account
	}
	if err := tmpl.Execute(w, data{Accounts: accounts}); err != nil {
		log.Fatalf("Error executing template: %v \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatalf("Error retrieving file: %v \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.OpenFile("/data/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error creating file: %v \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Fatalf("Error copying file: %v \n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

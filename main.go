package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml")
)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func executeTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /", homeHandler)
	r.HandleFunc("GET /contact", contactHandler)
	r.HandleFunc("GET /faq", faqHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

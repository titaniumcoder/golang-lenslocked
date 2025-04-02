package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/titaniumcoder/golang-lenslocked/controllers"
	"github.com/titaniumcoder/golang-lenslocked/templates"
	"github.com/titaniumcoder/golang-lenslocked/views"
)

func main() {
	r := http.NewServeMux()

	homeTpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	contactTpl := views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	faqTpl := views.Must(views.ParseFS(templates.FS, "faq.gohtml"))

	r.HandleFunc("GET /", controllers.StaticHandler(homeTpl))
	r.HandleFunc("GET /contact", controllers.StaticHandler(contactTpl))
	r.HandleFunc("GET /faq", controllers.StaticHandler(faqTpl))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))
}

package main

import (
	"fmt"
	"net/http"

	"github.com/titaniumcoder/golang-lenslocked/controllers"
	"github.com/titaniumcoder/golang-lenslocked/templates"
	"github.com/titaniumcoder/golang-lenslocked/views"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("GET /", controllers.StaticHandler(parsePage("home-page.html")))
	r.HandleFunc("GET /contact", controllers.StaticHandler(parsePage("contact-page.html")))
	r.HandleFunc("GET /faq", controllers.FAQ(parsePage("faq-page.html")))

	r.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(templates.StaticFS))))

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}

func parsePage(name string) views.Template {
	return views.Must(views.ParseFS(templates.FS, "tailwind.html", name))
}

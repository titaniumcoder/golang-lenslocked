package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/titaniumcoder/golang-lenslocked/controllers"
	"github.com/titaniumcoder/golang-lenslocked/templates"
	"github.com/titaniumcoder/golang-lenslocked/views"
)

func main() {
	fmt.Println(os.Getenv("DATABASE_URL"))
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!")
	defer db.Close()
	r := http.NewServeMux()

	r.HandleFunc("GET /", controllers.StaticHandler(parsePage("home-page.html")))
	r.HandleFunc("GET /contact", controllers.StaticHandler(parsePage("contact-page.html")))
	r.HandleFunc("GET /faq", controllers.FAQ(parsePage("faq-page.html")))

	var usersC controllers.Users
	usersC.Templates.New = parsePage("users/new.html")
	r.HandleFunc("GET /users/new", usersC.New)
	r.HandleFunc("POST /users", usersC.Create)

	r.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(templates.StaticFS))))

	fmt.Println("Starting the server on :3000...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}

func parsePage(name string) views.Template {
	return views.Must(views.ParseFS(templates.FS, "tailwind.html", name))
}

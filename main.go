package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/titaniumcoder/golang-lenslocked/controllers"
	"github.com/titaniumcoder/golang-lenslocked/models"
	"github.com/titaniumcoder/golang-lenslocked/templates"
	"github.com/titaniumcoder/golang-lenslocked/views"
)

func main() {
	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	if os.Getenv("CSRF_KEY") != "" {
		csrfKey = os.Getenv("CSRF_KEY")
	}
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO fix this before deploying
		csrf.Secure(false),
	)
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

	UserService := models.UserService{
		DB: db,
	}
	usersC := controllers.Users{
		UserService: &UserService,
	}
	usersC.Templates.New = parsePage("users/new.html")
	usersC.Templates.SignIn = parsePage("users/signin.html")
	r.HandleFunc("GET /users/signin", usersC.SignIn)
	r.HandleFunc("POST /users/signin", usersC.ProcessSignIn)
	r.HandleFunc("GET /users/new", usersC.New)
	r.HandleFunc("POST /users", usersC.Create)
	r.HandleFunc("GET /users/me", usersC.CurrentUser)

	r.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(templates.StaticFS))))

	fmt.Println("Starting the server on :3000...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, csrfMw(r))
}

func parsePage(name string) views.Template {
	return views.Must(views.ParseFS(templates.FS, "tailwind.html", name))
}

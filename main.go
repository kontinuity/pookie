package main

import (
	"html/template"
	"log"
	"net/http"
)

type LoginData struct {
	Title string
	Error string
}

func main() {
	http.HandleFunc("/", loginPageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	data := LoginData{
		Title: "Login - Pookie App",
		Error: "",
	}
	tmpl.Execute(w, data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "admin" && password == "password" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	data := LoginData{
		Title: "Login - Pookie App",
		Error: "Invalid username or password",
	}
	tmpl.Execute(w, data)
}
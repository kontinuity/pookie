package main

import (
	"html/template"
	"net/http"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("dashboard").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dashboard - Pookie App</title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0/css/all.min.css">
</head>
<body>
    <section class="hero is-success">
        <div class="hero-body">
            <div class="container has-text-centered">
                <h1 class="title">
                    <i class="fas fa-check-circle"></i>
                    Welcome to Dashboard!
                </h1>
                <h2 class="subtitle">
                    You have successfully logged in.
                </h2>
                <a class="button is-light" href="/">
                    <span class="icon">
                        <i class="fas fa-arrow-left"></i>
                    </span>
                    <span>Back to Login</span>
                </a>
            </div>
        </div>
    </section>
</body>
</html>
    `))
	tmpl.Execute(w, nil)
}
package handler

import (
	"html/template"
	"net/http"
)

// NotFoundHandler - Custom 404 handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/404.html"))

	w.WriteHeader(http.StatusNotFound) // Set the HTTP status code to 404
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not load the 404 page", http.StatusInternalServerError)
	}
}

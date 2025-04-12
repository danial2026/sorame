package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// ServeStatic serves static HTML, CSS, JS files from the "static" directory and handles 404 errors
func ServeStatic(staticDir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create the full path to the requested file
		filePath := filepath.Join(staticDir, r.URL.Path[len("/static/"):])

		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Serve the custom 404 page
			http.ServeFile(w, r, filepath.Join(staticDir, "404.html"))
			return
		}

		// If the file exists, serve it
		http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))).ServeHTTP(w, r)
	})
}

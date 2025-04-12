package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

// ServeStatic serves static HTML, CSS, JS files from the "static" directory and handles 404 errors
func ServeStatic(staticDir string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the path after removing the static prefix
		path := r.URL.Path
		if len(path) > 0 && path[0] == '/' {
			path = path[1:]
		}

		// Create the full path to the requested file
		filePath := filepath.Join(staticDir, path)

		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// Serve the custom 404 page
			http.ServeFile(w, r, filepath.Join(staticDir, "404.html"))
			return
		}

		// If the file exists, serve it
		fs.ServeHTTP(w, r)
	})
}

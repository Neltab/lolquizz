package http

import (
	"net/http"
	"os"
	"path/filepath"
)

// SPAHandler serves the React app and falls back to index.html
// for any path that isn't a real file (so TanStack Router handles it)
func SPAHandler(distPath string) http.HandlerFunc {
	fileServer := http.FileServer(http.Dir(distPath))

	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested file actually exists
		path := filepath.Join(distPath, r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// File doesn't exist — serve index.html (let React Router handle it)
			http.ServeFile(w, r, filepath.Join(distPath, "index.html"))
			return
		}
		// File exists — serve it directly (JS, CSS, images, etc.)
		fileServer.ServeHTTP(w, r)
	}
}

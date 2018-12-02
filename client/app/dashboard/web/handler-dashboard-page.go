package web

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func (wds *WebDashboardServer) serveDashboardPage(w http.ResponseWriter, r *http.Request) {
	templatePath := filepath.Join("www", "dashboard.html")

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(templatePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	keys := struct {
		MapAPIKey string
	}{
		MapAPIKey: wds.keyStorage.mapAPIKey,
	}

	if err := tmpl.Execute(w, keys); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

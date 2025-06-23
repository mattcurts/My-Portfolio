package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/mattcurts/portfolio/data"
)

func main() {
	// Parse all .html files under templates/
	tpl := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

	// Static file servers
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Main handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "base", data.Data); err != nil {
			log.Println("template error:", err)
		}
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

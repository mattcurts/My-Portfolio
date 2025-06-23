//go:generate go run main.go -static
package main

import (
	"flag"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mattcurts/portfolio/data"
)

var static = flag.Bool("static", false,
	"generate static site in docs/")

func main() {
	flag.Parse()

	tpl := template.Must(template.
		ParseGlob(filepath.Join("templates", "*.html")))

	if *static {
		out := "docs"
		os.RemoveAll(out)
		if err := os.MkdirAll(out, 0755); err != nil {
			log.Fatal(err)
		}

		// copy assets/
		if err := copyDir("assets", filepath.Join(out, "assets")); err != nil {
			log.Fatal(err)
		}

		// render base → docs/index.html
		f, err := os.Create(filepath.Join(out, "index.html"))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if err := tpl.ExecuteTemplate(f, "base", data.Data); err != nil {
			log.Fatal(err)
		}
		log.Println("Static site generated in", out)
		return
	}

	// dynamic server
	http.Handle("assets/",
		http.StripPrefix("assets/",
			http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "base", data.Data); err != nil {
			log.Println("template error:", err)
		}
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// copyDir recursively copies a src directory to dst.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if info.IsDir() {
			return os.MkdirAll(target, info.Mode())
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, info.Mode())
	})
}

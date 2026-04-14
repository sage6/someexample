package main

import (
	"embed"
	"html/template"
	"net/http"
	"io/fs"
)

//go:embed public/*
var staticFiles embed.FS

//go:embed *.gohtml
var tplFiles embed.FS

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFS(tplFiles, "*.gohtml"))
}

func main() {
	subFS, _ := fs.Sub(staticFiles, "public")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tpl.ExecuteTemplate(w, "index.gohtml", nil)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(subFS))))
	http.ListenAndServe(":8080", nil)
}

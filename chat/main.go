package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("chat/templates", t.filename)))
		//GOPATHからみてるのでそこからのパスを書く
	})
	t.templ.Execute(w, nil)
}

func main() {
	//route
	http.Handle("/", &templateHandler{filename: "chat.html"})
	//WebServer Starting...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

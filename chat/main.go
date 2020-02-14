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
	r := newRoom()
	http.Handle("/room", r)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	//start chatroom
	go r.run()
	//WebServer Starting...
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

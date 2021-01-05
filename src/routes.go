package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func routes() {
	// aqui solo es para dar la salida de informacion

	r := mux.NewRouter()
	// REDIRIJE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/1", 301)
		// yeah this is cool
	})
	// load the page
	r.HandleFunc("/{page:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "view/home.html")
	})
	// send the files
	r.HandleFunc("/public/{directory+}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, strings.ReplaceAll(r.URL.Path, "&", "/")[1:])
	})
	r.HandleFunc("/admin/postfile", newPost)
	// get the info
	r.HandleFunc("/api/{page:[0-9]+}", api)
	// render the publication

	r.HandleFunc("/publication/{id:[0-9]+}", renderInfo)

	log.Println(http.ListenAndServe(":8080", r))
}

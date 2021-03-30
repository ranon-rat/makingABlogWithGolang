package main

import (
	"log"
	"net/http"
	"os"

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
	r.HandleFunc(`/public/{file:[\/\w\d\W]+?}`, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	r.HandleFunc("/admin/postfile", newPost)
	// get the info
	r.HandleFunc("/api/{page:[0-9]+}", api)
	// render the publication

	r.HandleFunc("/publication/{id:[0-9]+}", renderInfo)
	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "8080"
	}

	log.Println(http.ListenAndServe(":"+port, r))
}

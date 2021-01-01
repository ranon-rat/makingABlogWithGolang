package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	// aqui solo es para dar la salida de informacion
	r := mux.NewRouter()
	r.HandleFunc("/admin/postfile", newPost)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.Redirect(w, r, "/0", 301)
		//http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc("/api/{page:[0-9]}", api)
	r.HandleFunc("/{page:[0-9]}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc("/publication/{id:[0-9]}", renderInfo)
	r.HandleFunc("/public/{directory}/{file}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, r.URL.Path[1:])

	})

	log.Println(http.ListenAndServe(":8080", r))
}

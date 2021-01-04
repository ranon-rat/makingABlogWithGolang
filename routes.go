package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func routes() {
	// aqui solo es para dar la salida de informacion
	log.Println(getTheSizeOfTheQuery())
	r := mux.NewRouter()
	// REDIRIJE
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number, err := getTheSizeOfTheQuery()
		if err != nil || number <= 0 {
			log.Println("something is wrong", err, number)
			w.Write([]byte("something is wrong sorry"))
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/%d", number/10), 301)
	})
	r.HandleFunc("/{page:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc("/public/{directory}/{file}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, r.URL.Path[1:])

	})
	r.HandleFunc("/admin/postfile", newPost)
	r.HandleFunc("/api/{page:[0-9]+}", api)

	r.HandleFunc("/publication/{id:[0-9]+}", renderInfo)

	log.Println(http.ListenAndServe(":8080", r))
}

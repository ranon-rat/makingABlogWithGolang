package main

import (
	"fmt"
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
		number, err := getTheSizeOfTheQuery()
		if err != nil || number <= 0 {
			log.Println("something is wrong", err, number)
			w.Write([]byte("something is wrong sorry"))
			return
		}

		// hmm por alguna razon no me redirect a la pagina que deseo hmm
		/*
			creo que ya se cual es el problema
			puede que sean por los datos de navegacion y que por
			defecto te reenvie a la pagina con la que accediste primero por alguna razon
			puede que sea buena idea  buscar una forma de hacer que borre las cookies
			si ya cheque , si es problema del navegador 😩. eso no se como solucionarlo :(

		*/
		// aqui lo que deberia de hacer es obtener el ultimo resultado y enviarlo

		http.Redirect(w, r, fmt.Sprintf("/%d", number/cantidad), 301)
	})
	r.HandleFunc("/{page:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "view/home.html")
	})
	r.HandleFunc("/public/{directory+}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, strings.ReplaceAll(r.URL.Path, "&", "/")[1:])
	})

	r.HandleFunc("/admin/postfile", newPost)
	r.HandleFunc("/api/{page:[0-9]+}", api)

	r.HandleFunc("/publication/{id:[0-9]+}", renderInfo)

	log.Println(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"
)

func routes() {
	// aqui solo es para dar la salida de informacion
	http.HandleFunc("/", renderInfo)
	http.HandleFunc("/public/", public)
	http.HandleFunc("/post", newPost)
	log.Println(http.ListenAndServe(":8080", nil))
}

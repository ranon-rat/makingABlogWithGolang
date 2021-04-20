package main

import (
	"log"

	"github.com/ranon-rat/blog/src/router"
)

// y el main obviamente
func main() {

	log.Println("starting server")
	router.Routes()

}

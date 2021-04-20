package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
	"github.com/ranon-rat/blog/src/dataControll"
	"github.com/ranon-rat/blog/src/stuff"
)

func RenderMarkdown(p chan stuff.Document, publicationChan chan stuff.Document, errChan chan error) {
	// lo que hace es parsear el markdown en html
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	// for now im doing this
	// but i want to use this with a db

	d, err := <-publicationChan, <-errChan
	if err != nil {
		p <- d
		log.Println(err)
		return
	}
	
	// ya sabe, concurrencia
	// obtiene el markdown
	d.Body = string(markdown.ToHTML([]byte(d.Body), parser, nil)) // despues lo pasa a html
	p <- d                                                        // al final hace lo siguiente

}
func RenderInfo(w http.ResponseWriter, r *http.Request) {
	attr := mux.Vars(r)
	// get the id of the publication
	id, err := strconv.Atoi(attr["id"])
	if err != nil {
		log.Println("fuck ", err)
		w.Write([]byte("sorry but agioue ´"))
		return
	}
	p := make(chan stuff.Document)
	// then decode the markdown to html
	d, errChan := make(chan stuff.Document), make(chan error)
	go dataControll.GetOnlyOnePublication(id, d, errChan)
	go RenderMarkdown(p, d, errChan)
	t, err := template.ParseFiles("view/template.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// the goroutines are the best
	//aqui estamos usando templates para evitar que tener que estar usando otra cosa
	t.Execute(w, <-p)
}
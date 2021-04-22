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
	"golang.org/x/sync/errgroup"
)

func RenderMarkdown(p chan stuff.Document, publicationChan chan stuff.Document) {
	// lo que hace es parsear el markdown en html
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	// for now im doing this
	// but i want to use this with a db

	d := <-publicationChan
	// ya sabe, concurrencia
	// obtiene el markdown
	d.Body = string(markdown.ToHTML([]byte(d.Body), parser, nil)) // despues lo pasa a html
	p <- d                                                        // al final hace lo siguiente

}
func RenderInfo(w http.ResponseWriter, r *http.Request) {
	attr := mux.Vars(r)
	var controlErrors errgroup.Group
	// get the id of the publication
	id, err := strconv.Atoi(attr["id"])
	if err != nil {
		log.Println("fuck ", err)
		w.Write([]byte("sorry but agioue ´"))
		return
	}
	p := make(chan stuff.Document)
	// then decode the markdown to html

	d := make(chan stuff.Document)
	controlErrors.Go(func() error {
		return dataControll.GetOnlyOnePublication(id, d)
	})
	go RenderMarkdown(p, d)
	
	t, _ := template.ParseFiles("view/template.html")
	
	if err=controlErrors.Wait();err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))

		return
	}
	// the goroutines are the best
	//aqui estamos usando templates para evitar que tener que estar usando otra cosa
	t.Execute(w, <-p)
}

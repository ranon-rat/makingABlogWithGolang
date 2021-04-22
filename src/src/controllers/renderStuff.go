package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
	"github.com/ranon-rat/blog/src/dataControll"
	"github.com/ranon-rat/blog/src/stuff"
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

	// get the id of the publication
	id,_ := strconv.Atoi(attr["id"])

	p := make(chan stuff.Document)
	// then decode the markdown to html

	d := make(chan stuff.Document)
	
	go dataControll.GetOnlyOnePublication(id, d)
	
	go RenderMarkdown(p, d)
	
	t, _ := template.ParseFiles("view/template.html")
	
	
	// the goroutines are the best
	//aqui estamos usando templates para evitar que tener que estar usando otra cosa
	t.Execute(w, <-p)
}

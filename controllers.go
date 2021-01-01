package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
)

func bodyRequest(r *http.Request) string {
	body, _ := ioutil.ReadAll(r.Body)
	// aqui lo que haces es pasar el body a un string para despues pasarlo a un json
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	newStr := buf.String()
	return newStr
}

func renderMarkdown(p chan document, id int) {
	// lo que hace es parsear el markdown en html
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	// for now im doing this
	// but i want to use this with a db
	var d publications
	d, err := getPublications(id, id)
	if err != nil || len(d.Publications) <= 0 {

		log.Println("something is wrong")

		p <- document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
		return
	}
	fmt.Println(d)
	// ya sabe, concurrencia
	// obtiene el markdown
	html := string(markdown.ToHTML([]byte(d.Publications[0].Body), parser, nil)) // despues lo pasa a html
	p <- document{Title: "hello world", Body: html}                              // al final hace lo siguiente

}
func renderInfo(w http.ResponseWriter, r *http.Request) {
	p := make(chan document)
	go renderMarkdown(p, 1)
	t, err := template.ParseFiles("view/template.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// the goroutines are the best
	//aqui estamos usando templates para evitar que tener que estar usando otra cosa
	t.Execute(w, <-p)
}

// this only is for the styles and script
func check(c chan bool, d document, w http.ResponseWriter) {

	_, err := http.Get(d.Mineatura)
	c <- d.Body == "" || d.Title == "" || d.Mineatura == "" || len(d.Body) >= 100000 || len(d.Title) >= 50 || len(d.Mineatura) >= 100 || err != nil

}

// this is the post manager , with this you can do really interesting things
func newPost(w http.ResponseWriter, r *http.Request) {
	// aqui solo es para ver los metodos
	switch r.Method {
	case "POST":
		// i need to do some data bases for do this
		fmt.Println("someone is trying to post something")

		var d document
		m := bodyRequest(r)
		json.Unmarshal([]byte(m), &d)
		cont := make(chan bool)
		go check(cont, d, w)
		if <-cont {
			fmt.Println("fuck")
			return
		}

		go addPublication(d)
		fmt.Println("yes")
		break
	case "GET":
		http.ServeFile(w, r, "view/post.html")
		w.Write([]byte("lmao you are trying to get the page"))
		break
	default:
		// solo acepta 2 metodos de request
		w.Write([]byte("ñao ñao voce es maricon"))
		break
	}
	fmt.Println("something")
}

// this is the api
func api(w http.ResponseWriter, r *http.Request) {
	// only send this
	min, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		log.Println("something is wrong")
		return
	}

	max := (min * 10) + 10

	var a publications

	a, _ = getPublications(min*10, max)

	b, err := json.Marshal(a)
	if err != nil {
		log.Println("something is wrong")
		return
	}

	w.Write(b)

}

// hello there
// this is only the setup

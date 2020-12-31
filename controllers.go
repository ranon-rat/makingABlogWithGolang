package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type document struct {
	Title     string `json:"title"`
	Mineatura string `json:"mineatura"`
	Body      string `json:"body"`
}

func bodyRequest(r *http.Request) string {
	body, _ := ioutil.ReadAll(r.Body)
	// aqui lo que haces es pasar el body a un string para despues pasarlo a un json
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	newStr := buf.String()
	return newStr
}

func openFile(file chan *os.File) {
	// en lo que hacemos la base de datos esto es lo que tenemos
	fs, err := os.Open("markdown.md")
	if err != nil {
		fmt.Println("fuck")
	}
	// lo envia a este canal para asi poder usar un modelo concurrente
	file <- fs
}

func renderMarkdown(p chan document) {
	// lo que hace es parsear el markdown en html
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	file := make(chan *os.File) // for now im doing this
	// but i want to use this with a db
	go openFile(file)                                // ya sabe, concurrencia
	md, _ := ioutil.ReadAll(<-file)                  // obtiene el markdown
	html := string(markdown.ToHTML(md, parser, nil)) // despues lo pasa a html
	p <- document{Title: "hello world", Body: html}  // al final hace lo siguiente

}
func renderInfo(w http.ResponseWriter, r *http.Request) {
	p := make(chan document)
	go renderMarkdown(p)
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
func public(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
	// este solo es para que puedan acceder a el public
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
		// if someone try to do something like send a void string the server response with this
		// esto lo que hace es verificar si no se exeden
		if d.Body == "" || d.Title == "" || d.Mineatura == "" {
			w.Write([]byte("what the fuck dont do that"))
			return
		} else if len(d.Body) >= 100000 {
			w.Write([]byte("wtf dont send that shit"))
			return
		} else if len(d.Title) >= 50 {
			w.Write([]byte("wtf dude , the title is sow fucking big"))
			return

		} else if len(d.Mineatura) >= 100 {
			w.Write([]byte("busca otra url mas pequeña maricon"))
			return
		}
		fmt.Println(d)
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

// hello there
// this is only the setup

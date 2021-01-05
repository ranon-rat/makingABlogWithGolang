package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

func renderMarkdown(p chan document, publicationChan chan document, errChan chan error) {
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
func renderInfo(w http.ResponseWriter, r *http.Request) {
	attr := mux.Vars(r)
	// get the id of the publication
	id, err := strconv.Atoi(attr["id"])
	if err != nil {
		log.Println("fuck ", err)
		w.Write([]byte("sorry but agioue ´"))
		return
	}
	p := make(chan document)
	// then decode the markdown to html
	d, errChan := make(chan document), make(chan error)
	go getOnlyOnePublication(id, d, errChan)
	go renderMarkdown(p, d, errChan)
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
	log.Println(d, err)
	c <- d.Body == "" || d.Title == "" || d.Mineatura == "" || len(d.Body) >= 100000 || len(d.Title) >= 50 || len(d.Mineatura) >= 100 || err != nil
}

// this is the post manager , with this you can do really interesting things
func newPost(w http.ResponseWriter, r *http.Request) {
	conf, err := ioutil.ReadFile("adminip.txt")
	if err != nil {
		log.Println(err)
		return
	}

	// aqui solo es para ver los metodos
	log.Println(r.Header.Get("x-forwarded-for"))
	if strings.Contains(string(conf), r.Header.Get("x-forwarded-for")) {
		switch r.Method {
		case "POST":
			// i need to do some data bases for do this

			// in the future i gona do something with this
			var d document
			m := bodyRequest(r)
			// decode the bodyRequest
			json.Unmarshal([]byte(m), &d)
			cont := make(chan bool)
			// this is for check if something is wrong
			go check(cont, d, w)
			if <-cont {
				log.Println("fuck")
				return
			}
			// add the publications
			go addPublication(d)

			break
		case "GET":
			http.ServeFile(w, r, "view/post.html")
			break
		default:
			// solo acepta 2 metodos de request
			w.Write([]byte("ñao ñao voce es maricon"))
			break
		}
	} else {
		http.ServeFile(w, r, "view/denegado.html")
	}

}

// this is the ap

func api(w http.ResponseWriter, r *http.Request) {
	// only send this
	// this is for use the apis
	min, err := strconv.Atoi(mux.Vars(r)["page"])
	if err != nil {
		w.Write([]byte("something is wrong"))
		return
	}

	// concurrency communication
	//the db management
	aChan, errChan := make(chan publications), make(chan error)

	// we use this function only one time so, im only usign a anon function 😩

	go getPublications(min, aChan, errChan)
	a, err := <-aChan, <-errChan
	a.Cantidad = cantidad
	if err != nil {
		log.Println("fuck", err)
	}
	// por alguna razon no me permitia obtener el tamño de esa manera 🤑
	a.Size, _ = getTheSizeOfTheQuery()

	b, err := json.Marshal(a)
	if err != nil {
		log.Println("something is wrong", err)
		return
	}

	// send the json
	w.Write(b)

}

// hello there
// this is only the setup

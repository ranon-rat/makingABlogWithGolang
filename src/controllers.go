package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
)

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
	c <- d.Body == "" || d.Title == "" || d.Mineatura == "" || err != nil
}
func encryptData(data string) string {
	sum := sha256.Sum256([]byte(data)) // we encript the data

	return hex.EncodeToString(sum[:])
}

// this is the post manager , with this you can do really interesting things
func newPost(w http.ResponseWriter, r *http.Request) {
	conf, err := ioutil.ReadFile("adminip.txt")
	if err != nil {
		log.Println(err)
		return
	}

	// aqui solo es para ver los metodos

	cookie, err := r.Cookie("ip")
	if err != nil {

		http.SetCookie(w, &http.Cookie{
			Name:    "ip",
			Value:   r.Header.Get("x-forwarded-for"),
			Expires: time.Now().AddDate(1, 0, 0),
		})
	}
	// new things
	// go back to the normal
	if strings.Contains(string(conf), encryptData(r.Header.Get("x-forwarded-for"))) || strings.Contains(string(conf), encryptData(cookie.Value)) {
		switch r.Method {
		case "POST":
			// i need to do some data bases for do this

			// in the future i gonna do something with this
			var d document

			// decode the bodyRequest
			json.NewDecoder(r.Body).Decode(&d)
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
	sizeChan, dChan := make(chan int), make(chan []document)

	// we use this function only one time so, im only usign a anon function 😩

	go getPublications(min, dChan)
	go getTheSizeOfTheQuery(sizeChan)
	api := publications{
		Cantidad:     cantidad,
		Publications: <-dChan,
		Size:         <-sizeChan,
	}

	// send the json
	json.NewEncoder(w).Encode(api)

}

// hello there
// this is only the setup

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ranon-rat/blog/src/dataControll"
	"github.com/ranon-rat/blog/src/stuff"
)

// this is the post manager , with this you can do really interesting things
func NewPost(w http.ResponseWriter, r *http.Request) {
	confB, err := ioutil.ReadFile("database/adminip.txt")

	conf:=string(confB)
	// aqui solo es para ver los metodos

	cookie, err := r.Cookie("ip")
	if err != nil {
		cookie= &http.Cookie{
			Name:    "ip",
			Value:   r.Header.Get("x-forwarded-for"),
			Expires: time.Now().AddDate(1, 0, 0),
		}
		http.SetCookie(w,cookie)
		
	}

	// new things
	// go back to the normal
	if strings.Contains(conf, stuff.EncryptData(r.Header.Get("x-forwarded-for"))) || strings.Contains(conf, stuff.EncryptData(cookie.Value)) {
		switch r.Method {
		case "POST":
			// i need to do some data bases for do this

			// in the future i gonna do something with this
			var d stuff.Document

			// decode the bodyRequest
			json.NewDecoder(r.Body).Decode(&d)
			cont := make(chan bool)
			// this is for check if something is wrong
			go Check(cont, d, w)
			if <-cont {
				log.Println("fuck")
				return
			}
			// add the publications
			go dataControll.AddPublication(d)
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
		//log.Println(conf)
		http.ServeFile(w, r, "view/denegado.html")
	}

}
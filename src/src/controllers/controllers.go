package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ranon-rat/blog/src/dataControll"
	"github.com/ranon-rat/blog/src/stuff"
)

// this only is for the styles and script
func Check(c chan bool, d stuff.Document) {
	_, err := http.Get(d.Mineatura)
	log.Println(d, err)
	c <- d.Body == "" || d.Title == "" || d.Mineatura == "" || err != nil
}

// this is the ap

func Api(w http.ResponseWriter, r *http.Request) {
	// only send this
	// this is for use the apis
	min, _ := strconv.Atoi(mux.Vars(r)["page"])


	// concurrency communication
	//the db management
	sizeChan, dChan := make(chan int), make(chan []stuff.Document)

	// we use this function only one time so, im only usign a anon function 😩

	go dataControll.GetPublications(min, dChan)

	go dataControll.GetTheSizeOfTheQuery(sizeChan)

	
	api := stuff.Publications{
		Cantidad:     stuff.Cantidad,
		Publications: <-dChan,
		Size:         <-sizeChan,
	}
	// send the json
	json.NewEncoder(w).Encode(api)

}

// hello there
// this is only the setup

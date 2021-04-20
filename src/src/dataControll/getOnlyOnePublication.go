package dataControll

import (
	"log"

	"github.com/ranon-rat/blog/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)



func GetOnlyOnePublication(id int, aChan chan stuff.Document, errChan chan error) {
	q := (`SELECT * FROM publ
	WHERE id=?1`)
	db := GetConnection()
	defer db.Close()
	p, err := db.Query(q,id)
	if err != nil {
		log.Println("something is wrong", err)
		aChan <- stuff.Document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
		errChan <- err
		return
	}

	defer p.Close()
	var d stuff.Document
	for p.Next() {
		/*
		(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    mineatura TEXT NOT NULL,
    body TEXT NOT NULL
		*/
		err = p.Scan(&d.ID, &d.Title, &d.Mineatura, &d.Body)
		if err != nil {
			log.Println(err)
			aChan <- stuff.Document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
			errChan <- err
			return
		}
	}
	aChan <- d
	errChan <- nil
}
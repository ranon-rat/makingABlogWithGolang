package dataControll

import (
	"github.com/ranon-rat/blog/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)

func GetOnlyOnePublication(id int, aChan chan stuff.Document) error {
	q := (`SELECT * FROM publ WHERE id=?1`)
	db := GetConnection()
	defer db.Close()
	p, err := db.Query(q, id)
	if err != nil {

		aChan <- stuff.Document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
		return err
	}

	defer p.Close()
	var d stuff.Document
	for p.Next() {

		err = p.Scan(&d.ID, &d.Title, &d.Mineatura, &d.Body)
		if err != nil {

			aChan <- stuff.Document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
			return err
		}
	}
	aChan <- d
	return nil
}

package dataControll

import (
	"errors"
	"log"

	"github.com/ranon-rat/blog/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)


func AddPublication(e stuff.Document) error {
	q := `INSERT INTO 
	publ(title,mineatura,body) 
	values($1,$2,$3);
	`
	db := GetConnection()
	defer db.Close()
	stm, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stm.Close()
	r, err := stm.Exec(&e.Title, &e.Mineatura, &e.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		log.Println("se esperaba una sola fila omg")
		return errors.New("se esperaba una sola fila omg")
	}
	return nil
}
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type publication struct {
	id        int
	titulo    string
	mineatura string
	body      string
}

func getConnection() *sql.DB {
	dsn := "postgres://ranon:ranon@127.0.0.1:5432/publications?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}
	return db
	//db.Query("SELECT * FROM publications WHERE id <=10 AND id>0;")
}
func getPublications(min, max int) {
	q := fmt.Sprintf("SELECT * FROM publ WHERE id <=%d AND id>%d", max, min)
	db := getConnection()
	defer db.Close()
	m, err := db.Query(q)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer m.Close()
	for m.Next() {
		var p publication
		err := m.Scan(&p.id, &p.titulo, &p.mineatura, &p.body)
		if err != nil {
			fmt.Println("fuck", err)
		}
		fmt.Println(p)
	}

}
func addPublication(e publication) error {
	q := `INSERT INTO 
	publ(titulo,mineatura,body) 
	values($1,$2,$3)
	
	`
	db := getConnection()
	defer db.Close()
	stm, err := db.Prepare(q)
	if err != nil {
		log.Println(err)
		return err

	}
	defer stm.Close()
	r, err := stm.Exec(&e.titulo, &e.mineatura, &e.body)
	if err != nil {
		log.Println(err)
		return err
	}
	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("se esperaba una sola fila omg")
	}
	return nil

}
func main() {

	pub := publication{
		titulo:    "yes",
		mineatura: "yes",
		body:      "oh yes",
	}
	addPublication(pub)
	getPublications(-10, 10)

	fmt.Println("hello world")
}

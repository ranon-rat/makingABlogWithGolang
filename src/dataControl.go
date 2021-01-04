package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type document struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Mineatura string `json:"mineatura"`
	Body      string `json:"bodyOfDocument"`
}
type publications struct {
	Size         int        `json:"Size"`
	Publications []document `json:"Publications"`
	Cantidad     int        `json:"Cantidad"`
}

func getConnection() *sql.DB {
	dsn := "postgres://ranon:ranon@127.0.0.1:5432/publications?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func addPublication(e document) error {
	q := `INSERT INTO 
	publ(titulo,mineatura,body) 
	values($1,$2,$3);
	`
	db := getConnection()
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
		return errors.New("se esperaba una sola fila omg")
	}
	return nil
}
func getPublications(min int, pChan chan publications, errChan chan error) {
	// este es el consultorio croe que se llamaba asi , ya no me acuerdo xd
	q := fmt.Sprintf(`
	SELECT * FROM publ 
	WHERE  id >=%d
	ORDER BY id DESC 
	LIMIT %d;`, min, cantidad)
	/*
		aqui lo que basicamente hace es ordenar del mayor al menor
	*/
	// aqui lo que hace es ordenar el resultado
	db := getConnection()
	// aqui lo que hace es conectarse a la base de datos
	defer db.Close()
	//espera a cerrarse para evitar ciertos problemas de seguridad
	m, err := db.Query(q) // envia esto y la salida deb de ser la siguiente
	if err != nil {

		pChan <- publications{}
		errChan <- err
		log.Println(err)
		return
	}
	defer m.Close() // espera a cerrar el canal ( por razones de seguridad)

	var pubs publications
	for m.Next() {
		// repasa la informacion,
		var d document
		// cambia los valores de publication
		err := m.Scan(&d.ID, &d.Title, &d.Mineatura, &d.Body)
		if err != nil {
			pChan <- publications{}
			errChan <- err
			fmt.Println(err)
			return
		}
		pubs.Publications = append(pubs.Publications, d)
		// los agrega a una listaa
	}

	pChan <- pubs
	errChan <- nil

	return
}
func getOnlyOnePublication(id int, aChan chan document, errChan chan error) {
	q := fmt.Sprintf(`
	SELECT * FROM publ
	WHERE id=%d`, id)
	db := getConnection()
	defer db.Close()
	p, err := db.Query(q)
	if err != nil {

		log.Println("something is wrong", err)
		aChan <- document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
		errChan <- err
		return
	}

	defer p.Close()
	var d document
	for p.Next() {
		err = p.Scan(&d.ID, &d.Title, &d.Mineatura, &d.Body)
		if err != nil {
			log.Println(err)
			aChan <- document{Title: "sorry but something is wrong", Body: "<h1> something wrong </h1>"}
			errChan <- err
			return
		}
	}
	aChan <- d
	errChan <- nil

}

// this is for get the size of the table
func getTheSizeOfTheQuery() (int, error) {
	q := `
	SELECT MAX(id) 
	FROM publ
	`
	// como no he encontrado muchas maneras de encontrar el
	//tamaño de una tabla lo que hace aqui es basicamente seleccionar el maximo valor

	var dataSize int
	db := getConnection()
	defer db.Close()
	m, err := db.Query(q)
	if err != nil {
		return 0, err
	}
	defer m.Close()
	for m.Next() {

		err = m.Scan(&dataSize)
		if err != nil {
			return 0, err
		}

	}
	return dataSize, nil
}

package main

import (
	"database/sql"
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

// aqui lo que estoy haciendo es conectarme a la base de datos
func getConnection() *sql.DB {
	dsn := "postgres://ranon:ranon@127.0.0.1:5432/publications?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func getPublications(min, max int) []publication {
	// este es el consultorio croe que se llamaba asi , ya no me acuerdo xd
	q := fmt.Sprintf("SELECT * FROM publ WHERE id <=%d AND id>%d", max, min)
	db := getConnection()
	// aqui lo que hace es conectarse a la base de datos
	defer db.Close()
	//espera a cerrarse para evitar ciertos problemas de seguridad
	m, err := db.Query(q) // envia esto y la salida deb de ser la siguiente
	if err != nil {
		fmt.Println(err) // solo por si hay un error xd
		return nil
	}
	defer m.Close() // espera a cerrar el canal ( por razones de seguridad)

	var pubs []publication
	for m.Next() {
		// repasa la informacion,
		var p publication
		// cambia los valores de publication
		err := m.Scan(&p.id, &p.titulo, &p.mineatura, &p.body)
		if err != nil {
			// en caso de que haya un error
			log.Println("fuck", err)
		}
		pubs = append(pubs, p)
		// los agrega a una listaa
	}
	return pubs
}

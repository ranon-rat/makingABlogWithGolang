package dataControll

import (
	"github.com/ranon-rat/blog/src/stuff"

	_ "github.com/mattn/go-sqlite3"
)

func GetPublications(min int, pChan chan []stuff.Document) {
	sChan := make(chan int)
	go GetTheSizeOfTheQuery(sChan)
	size := <-sChan
	// este es el consultorio croe que se llamaba asi , ya no me acuerdo xd
	q := `
	SELECT * FROM publ 
	WHERE  rowid >=?1 AND  rowid <=?2
	ORDER BY id DESC ;`

	/*
		aqui lo que basicamente hace es ordenar del mayor al menor
	*/
	// aqui lo que hace es ordenar el resultado
	db := GetConnection()
	// aqui lo que hace es conectarse a la base de datos
	defer db.Close()
	//espera a cerrarse para evitar ciertos problemas de seguridad
	m, _:= db.Query(q, (size - (min * stuff.Cantidad)), (size-(min*stuff.Cantidad)+stuff.Cantidad)+1) // envia esto y la salida deb de ser la siguiente
	
	defer m.Close() // espera a cerrar el canal ( por razones de seguridad)

	var pubs []stuff.Document
	for m.Next() {
		// repasa la informacion,
		var d stuff.Document
		// cambia los valores de publication
		m.Scan(&d.ID, &d.Title, &d.Mineatura, &d.Body)

		pubs = append(pubs, d)
		// los agrega a una listaa
	}

	pChan <- pubs


}

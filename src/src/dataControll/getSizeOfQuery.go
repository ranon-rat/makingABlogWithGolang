package dataControll

import (
	_ "github.com/mattn/go-sqlite3"
)

func GetTheSizeOfTheQuery(sizeChan chan int) {
	q := `SELECT COUNT(*) FROM publ`
	// como no he encontrado muchas maneras de encontrar el
	//tamaño de una tabla lo que hace aqui es basicamente seleccionar el maximo valor
	dataSize:=0
	db := GetConnection()
	defer db.Close()
	m, _ := db.Query(q)
	defer m.Close()
	for m.Next() {
		m.Scan(&dataSize)
	}
	return 
}
package dataControll

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetTheSizeOfTheQuery(sizeChan chan int){
	q := `SELECT COUNT(*) FROM publ`
	// como no he encontrado muchas maneras de encontrar el
	//tamaño de una tabla lo que hace aqui es basicamente seleccionar el maximo valor
	dataSize:=0
	db := GetConnection()
	defer db.Close()
	m, err := db.Query(q)
	if err!=nil{
		close(sizeChan)
		return 
	}
	
	for m.Next() {
		if err=m.Scan(&dataSize);err!=nil{
			log.Println(err)
			close(sizeChan)
			return 

		}
	}
	sizeChan<-dataSize

}

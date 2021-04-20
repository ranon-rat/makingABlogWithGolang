package dataControll

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("sqlite3", "./database/publications.db")

	if err != nil {
		fmt.Println(err)
	}
	return db
}



// this is for get the size of the table

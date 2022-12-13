package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Shelf struct {
	ID   int
	Name string
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "data.db")

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}

func GetShelfs() {

	var shelfs Shelf

	db, errDB := Connect()

	defer db.Close()

	if errDB != nil {
		log.Fatalln("Can't connect!")
	}

	rows, err := db.Query("SELECT * FROM shelf")

	if err != nil {
		fmt.Println("Não achei shelfs")
	}

	for rows.Next() {
		err := rows.Scan(&shelfs.ID, &shelfs.Name)
		if err != nil {
			fmt.Println("Não consegui adicionar shelf")
		}
	}

	if err != nil {
		log.Println("Sem data")
	}

	fmt.Println(shelfs)
}

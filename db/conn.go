package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "SQL@123456sql"
	dbname   = "postgres"
)

func ConncetDB () (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ 
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectato a: " + dbname)
	return db, nil

}












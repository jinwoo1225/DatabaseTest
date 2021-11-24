package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {

	DbUrl := os.Getenv("DB_URL")
	DbUser := os.Getenv("DB_USER")
	DbPass := os.Getenv("DB_PASS")
	//DbDatabase := os.Getenv("DB_DATABASE")

	db, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/", DbUser, DbPass, DbUrl))
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("CONNECTED!!")

	defer db.Close()
}

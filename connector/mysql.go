package connector

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	cfg := mysql.Config{
		User: "root",
		Passwd: "",
		Net: "tcp",
		Addr: "127.0.0.1:3306",
		DBName: "golang_demo",
		AllowNativePasswords: true,
	}

	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database !!!")
}

func Query(query string) *sql.Rows {
	rows, err := Db.Query(query)

	if err != nil {
		log.Fatal("Error when execute query.")
	}

	return rows
}

func QueryRow(query string) *sql.Row {
	return Db.QueryRow(query)
}
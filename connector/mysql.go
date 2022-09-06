package connector

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Db *sql.DB

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error when loading .env")
	}
	cfg := mysql.Config{
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net: os.Getenv("DB_NET"),
		Addr: os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

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
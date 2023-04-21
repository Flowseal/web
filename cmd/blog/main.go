package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const port = ":3000"

func openDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "root:root@tcp(localhost:3306)/blog")
}

func main() {
	db, err := openDatabase()
	if err != nil {
		log.Fatal(err)
	}	

	dbx := sqlx.NewDb(db, "mysql")

	mux := http.NewServeMux()
	mux.HandleFunc("/", catchAllHandler(dbx))

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Println("Starting server at http://localhost" + port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
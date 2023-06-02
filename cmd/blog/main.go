package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
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

	mux := mux.NewRouter()
	mux.HandleFunc("/", catchAllHandler(dbx))
	mux.HandleFunc("/home", catchAllHandler(dbx))	
	mux.HandleFunc("/index", catchAllHandler(dbx))
	mux.HandleFunc("/admin", catchAllHandler(dbx))	
	mux.HandleFunc("/login", catchAllHandler(dbx))	
	mux.HandleFunc("/post/{postId}/{postTitle}", catchAllHandler(dbx))

	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	
	log.Println("Starting server at http://localhost" + port)

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
package main

import (
	"os"
	"database/sql"
	"log"
	"net/http"
	"fmt"
)

var db *sql.DB

func init() {
	fmt.Println("Establishing connection ...")
	tmpDB, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=books_database sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		fmt.Println("error in sql.open")
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListBooks)
	http.HandleFunc("/book.html", handleViewBook)
	http.HandleFunc("/save", handleSaveBook)
	http.HandleFunc("/delete", handleDeleteBook)
	http.HandleFunc("/check", check)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

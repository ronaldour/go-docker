package main

import (
	"os"
	"database/sql"
	"log"
	"net/http"
	"fmt"
	"github.com/lib/pq"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var db *sql.DB

func init() {
	// fmt.Println("Establishing connection ...")
	// tmpDB, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=books_database sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")))
	// if err != nil {
	// 	fmt.Println("error in sql.open")
	// 	log.Fatal(err)
	// }
	// db = tmpDB
	fmt.Println("Establishing connection ...")
	sqltrace.Register("postgres", &pq.Driver{})
	tmpDB, err := sqltrace.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=books_database sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD")))
	if err != nil {
		fmt.Println("error in sql.open")
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	addr := os.Getenv("DD_AGENT_HOST") + ":" + os.Getenv("DD_TRACE_AGENT_PORT")
	fmt.Println("Host addr: " + addr)
	tracer.Start(tracer.WithAgentAddr(addr))
	defer tracer.Stop()

	mux := httptrace.NewServeMux(httptrace.WithServiceName("l2-demo-app")) // init the http tracer

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	mux.HandleFunc("/", handleListBooks)
	mux.HandleFunc("/book.html", handleViewBook)
	mux.HandleFunc("/save", handleSaveBook)
	mux.HandleFunc("/delete", handleDeleteBook)
	mux.HandleFunc("/check", check)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver

	"github.com/gorilla/handlers"

	"github.com/dmkales/blog-api/cmd/data"
	"github.com/dmkales/blog-api/internal/pkg/http/rest"
)

var db *sql.DB

func main() {
	port := ":8080"

	conn, err := dbconn.Setup()
	if err != nil {
		panic(err)
	}

	db = conn

	router := rest.Handler(db)

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})

	fmt.Printf("Running Server! on Port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

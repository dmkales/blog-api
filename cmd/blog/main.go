package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/dmkales/blog-api/internal/app/blog"
)

var db *sql.DB

func main() {
	port := ":8080"

	conn, err := dbinterface.Setup()
	if err != nil {
		panic(err)
	}

	db = conn

	var router = mux.NewRouter()
	router.HandleFunc("/task/list", getAllTasks).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})

	fmt.Printf("Running Server! on Port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := dbinterface.GetAllTask(db)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(tasks)
}

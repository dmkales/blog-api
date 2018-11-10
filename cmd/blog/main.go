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
	router.HandleFunc("/user", insertUser).Methods("POST")
	router.HandleFunc("/user/list", getAllUser).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")

	headersOk := handlers.AllowedHeaders([]string{"Authorization", "X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "DELETE", "POST", "PUT", "OPTIONS"})

	fmt.Printf("Running Server! on Port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := dbinterface.GetAll(db)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, err := dbinterface.Get(db, id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(user)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var user dbinterface.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		panic(err)
	}
	if err := dbinterface.Create(db, user); err != nil {
		panic(err)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user dbinterface.User
	params := mux.Vars(r)
	id := params["id"]
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		panic(err)
	}
	if err := dbinterface.Update(db, id, user); err != nil {
		panic(err)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if err := dbinterface.Delete(db, id); err != nil {
		panic(err)
	}
}

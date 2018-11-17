package user

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllUser : GetAllUser
func GetAllUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetAll(db)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(users)
	}
}

// GetUser : GetUser
func GetUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["user_id"]
		user, err := Get(db, id)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

// CreateUser : CreateUser
func CreateUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			panic(err)
		}
		if err := Create(db, user); err != nil {
			panic(err)
		}
	}
}

// UpdateUser : UpdateUser
func UpdateUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		params := mux.Vars(r)
		id := params["user_id"]
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			panic(err)
		}
		if err := Update(db, id, user); err != nil {
			panic(err)
		}
	}
}

// DeleteUser : DeleteUser
func DeleteUser(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["user_id"]
		if err := Delete(db, id); err != nil {
			panic(err)
		}
	}
}

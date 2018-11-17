package post

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllPost : GetAllPost
func GetAllPost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := GetAll(db)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(posts)
	}
}

// GetUserPost : GetUserPost
func GetUserPost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["user_id"]
		post, err := GetAllUserPost(db, id)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(post)
	}
}

// GetPost : GetPost
func GetPost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["post_id"]
		post, err := Get(db, id)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(post)
	}
}

// CreatePost : CreatePost
func CreatePost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&post); err != nil {
			panic(err)
		}
		if err := Create(db, post); err != nil {
			panic(err)
		}
	}
}

// UpdatePost : UpdatePost
func UpdatePost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		params := mux.Vars(r)
		id := params["post_id"]
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&post); err != nil {
			panic(err)
		}
		if err := Update(db, id, post); err != nil {
			panic(err)
		}
	}
}

// DeletePost : DeletePost
func DeletePost(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["post_id"]
		if err := Delete(db, id); err != nil {
			panic(err)
		}
	}
}

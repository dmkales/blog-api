package rest

import (
	"database/sql"
	"net/http"

	"github.com/dmkales/blog-api/internal/pkg/user"
	"github.com/gorilla/mux"
)

// Handler : returns all routes
func Handler(db *sql.DB) http.Handler {
	var router = mux.NewRouter()

	router.HandleFunc("/user", user.CreateUser(db)).Methods("POST")
	router.HandleFunc("/user/list", user.GetAllUser(db)).Methods("GET")
	router.HandleFunc("/user/{id}", user.GetUser(db)).Methods("GET")
	router.HandleFunc("/user/{id}", user.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/user/{id}", user.DeleteUser(db)).Methods("DELETE")

	return router
}

package rest

import (
	"database/sql"
	"net/http"

	"github.com/dmkales/blog-api/internal/pkg/post"
	"github.com/dmkales/blog-api/internal/pkg/user"
	"github.com/gorilla/mux"
)

// Handler : returns all routes
func Handler(db *sql.DB) http.Handler {
	var router = mux.NewRouter()

	router.HandleFunc("/user", user.CreateUser(db)).Methods("POST")
	router.HandleFunc("/user/list", user.GetAllUser(db)).Methods("GET")
	router.HandleFunc("/user/{user_id}", user.GetUser(db)).Methods("GET")
	router.HandleFunc("/user/{user_id}", user.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/user/{user_id}", user.DeleteUser(db)).Methods("DELETE")

	router.HandleFunc("/post", post.CreatePost(db)).Methods("POST")
	router.HandleFunc("/post/list/{user_id}", post.GetUserPost(db)).Methods("GET")
	router.HandleFunc("/post/list", post.GetAllPost(db)).Methods("GET")
	router.HandleFunc("/post/{post_id}", post.GetPost(db)).Methods("GET")
	router.HandleFunc("/post/{post_id}", post.UpdatePost(db)).Methods("PUT")
	router.HandleFunc("/post/{post_id}", post.DeletePost(db)).Methods("DELETE")

	return router
}

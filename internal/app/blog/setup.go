package dbinterface

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// User : structure for inserting and displaying user fields
type User struct {
	UserID    int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	Created   *time.Time `json:"created"`
	Updated   *time.Time `json:"updated"`
}

// Setup : initialize the db connection and returns the instance
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/blog?parseTime=true", "chael", "password"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

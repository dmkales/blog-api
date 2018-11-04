package dbinterface

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Task : structure for inserting and displaying task fields
type Task struct {
	ID      int        `json:"id"`
	Task    string     `json:"task"`
	IsDone  bool       `json:"is_done"`
	Created *time.Time `json:"created"`
}

// Setup : initialize the db connection and returns the instance
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/gocookbook?parseTime=true", "chael", "password"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

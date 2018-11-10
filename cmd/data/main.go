package dbconn

import (
	"database/sql"
	"fmt"
)

// Setup : initialize the db connection and returns the instance
func Setup() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/blog?parseTime=true", "chael", "password"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

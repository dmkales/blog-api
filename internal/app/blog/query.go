package dbinterface

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// GetAll : get all task
func GetAll(db *sql.DB) ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT user_id, first_name, last_name, email, created, updated FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Created, &u.Updated); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// Get : get a single user using and ID
func Get(db *sql.DB, id string) (User, error) {
	var u User
	rows, err := db.Query("SELECT user_id, first_name, last_name, email, created, updated FROM user WHERE user_id=?", id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Created, &u.Updated); err != nil {
			return User{}, err
		}
	}
	return u, rows.Err()
}

// Create : insert a user into the user table
func Create(db *sql.DB, user User) error {
	sqlStmt := `INSERT INTO user(first_name, last_name, email, password, created) VALUES (?, ?, ?, ?, NOW())`
	if _, err := db.Exec(sqlStmt, user.FirstName, user.LastName, user.Email, user.Password); err != nil {
		return err
	}
	return nil
}

// Update : update a single user
func Update(db *sql.DB, id string, user User) error {
	return nil
}

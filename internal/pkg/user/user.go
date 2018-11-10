package user

import "time"

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

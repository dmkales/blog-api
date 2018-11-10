package post

import "time"

// Post : structuer for inserting and displaying post fields
type Post struct {
	PostID  int        `json:"id"`
	UserID  int        `json:"user_id"`
	Title   string     `json:"title"`
	Body    string     `json:"body"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

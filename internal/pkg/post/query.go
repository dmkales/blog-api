package post

import (
	"database/sql"
	"time"
)

// GetAll : get all posts
func GetAll(db *sql.DB) ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT post_id, user_id, title, body, created, updated FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Body, &p.Created, &p.Updated); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// GetAllUserPost : get all user post using and ID
func GetAllUserPost(db *sql.DB, id string) ([]Post, error) {
	var posts []Post
	rows, err := db.Query("SELECT post_id, user_id, title, body, created, updated FROM post WHERE user_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Body, &p.Created, &p.Updated); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

// Get : get a single post using and Post ID
func Get(db *sql.DB, id string) (Post, error) {
	var p Post
	rows, err := db.Query("SELECT post_id, user_id, title, body, created, updated FROM post WHERE post_id=?", id)
	if err != nil {
		return Post{}, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&p.PostID, &p.UserID, &p.Title, &p.Body, &p.Created, &p.Updated); err != nil {
			return Post{}, err
		}
	}
	return p, rows.Err()
}

// Create : insert a post into the post table
func Create(db *sql.DB, p Post) error {
	sqlStmt := `INSERT INTO post(user_id, title, body, created) VALUES (?, ?, ?, NOW())`
	if _, err := db.Exec(sqlStmt, p.UserID, p.Title, p.Body); err != nil {
		return err
	}
	return nil
}

// Update : update a user's post
func Update(db *sql.DB, id string, p Post) error {
	t := time.Now()
	tf := t.Format("2006-01-02 15:04:05")
	sqlStmt := `UPDATE post SET title = ?, body = ?, updated = ? WHERE post_id = ?`
	if _, err := db.Exec(sqlStmt, p.Title, p.Body, tf, id); err != nil {
		return err
	}
	return nil
}

// Delete : delete a post of a user
func Delete(db *sql.DB, id string) error {
	sqlStmt := `DELETE FROM post WHERE post_id = ?`
	if _, err := db.Exec(sqlStmt, id); err != nil {
		return err
	}
	return nil
}

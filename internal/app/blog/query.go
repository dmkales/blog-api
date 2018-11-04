package dbinterface

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// GetAllTask : get all task
func GetAllTask(db *sql.DB) ([]Task, error) {
	var tasks []Task
	rows, err := db.Query("SELECT id, task, is_done, created FROM example")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Task, &t.IsDone, &t.Created); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

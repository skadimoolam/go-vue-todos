package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done string `json:"done"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks;"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Done)
		if err != nil {
			panic(err)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PostTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name, done) VALUES(?, 'false')"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func PutTask(db *sql.DB, task Task) (int64, error) {
	var sql string

	sql = "UPDATE tasks SET name = ?, done = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(task.Name, task.Done, task.ID)

	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

package main

import (
	"database/sql"

	"./handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := initDb("storage.db")
	migrate(db)

	e := echo.New()
	e.Static("/", "public")
	e.Get("/tasks", handlers.GetTasks(db))
	e.POST("/task", handlers.PostTask(db))
	e.Put("/task", handlers.PutTask(db))
	e.Delete("/task/:id", handlers.DeleteTask(db))
	e.Run(standard.New(":8080"))
}

func initDb(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			done INTEGER NOT NULL
		);
	`

	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

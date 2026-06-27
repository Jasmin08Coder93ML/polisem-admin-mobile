package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// NewPostgresDB - DB baglanyşygyny döredýär
func NewPostgresDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	// DB-niň işläp durandygyny barlamak
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database birikmesi üstünlikli!")
	return db, nil
}

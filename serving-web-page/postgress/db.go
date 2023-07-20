package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// gobal DB variable
var Connection *sql.DB

// initialDB creates a new instanmce of DB
func InitDB() {
	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		"postgres", "mysecretpassword", "localhost", 5432, "postgres",
	)

	var err error
	Connection, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
}

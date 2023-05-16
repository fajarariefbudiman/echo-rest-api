package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewData() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/echo_rest")
	if err != nil {
		panic(err)
	}

	return db
}

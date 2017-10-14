package Repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
)

const DB_CONNECT_STRING = "host=localhost port=5432 user=tuser password=123 dbname=atend sslmode=disable"

type Postgre struct {
	sql.DB
}

func (this *Postgre) Open() *sql.DB {
	db, err := sql.Open("postgres", DB_CONNECT_STRING)
	//defer db.Close()

	if err != nil {
		// Пишем в лог и паникуем
		panic("Database error")
	}

	return db
}

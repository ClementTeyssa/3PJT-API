package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	//TODO: set var by .env
	db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=localhost port=5432 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
}

	// Create Table users if not exists
func createUsersTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
	if err != nil {
		log.Fatal(err)
	}
}

// Getter for db var
func GetDb() *sql.DB {
	return db
}
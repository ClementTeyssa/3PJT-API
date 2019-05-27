package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func DatabaseInit() {
	var err error
	//TODO: set var by .env
	db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=postgres port=5432 sslmode=disable")
	// db, err = sql.Open("postgres", "user=test dbname=goapi password=test host=localhost port=5432 sslmode=disable")

	if err != nil {
		log.Panic(err)
	}
	createUsersTable()
	createTransactionsTable()
	createBlocksTable()
}

// Create Table users if not exists
func createUsersTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users(id serial, email varchar(100), password varchar, adress varchar, privatekey bytea, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
	if err != nil {
		log.Panic(err)
	}
}

func createTransactionsTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS transactions(id serial, accountfrom varchar, accountto varchar, amount float, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
	if err != nil {
		log.Panic(err)
	}
}

func createBlocksTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS blocks(id serial, timestamp int, transactionid int, hash varchar, prevhash varchar, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")
	if err != nil {
		log.Panic(err)
	}
}

// Getter for db var
func GetDb() *sql.DB {
	return db
}

package models

import (
	"github.com/ClementTeyssa/New_Test/config"
	"log"
	"time"
)

type User struct {
	Id           int       `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Design       string    `json:"design"`
	Style        string    `json:"style"`
	Doors        uint8     `json:"doors"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Users []User

func NewUser(user *User) {
	if user == nil {
		log.Fatal(user)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO users (manufacturer, design, style, doors, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;", user.Manufacturer, user.Design, user.Style, user.Doors, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	
	if err != nil {
		log.Fatal(err)
	}
}

func FindUserById(id int) *User {
	var user User
	row := config.GetDb().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.Id, &user.Manufacturer, &user.Design, &user.Style, &user.Doors, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func AllUsers() *Users {
	var users Users
	rows, err := config.GetDb().Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	// Close rows after all readed
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Manufacturer, &user.Design, &user.Style, &user.Doors, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return &users
}

func UpdateUser(user *User) {
	user.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE users SET manufacturer=$1, design=$2, style=$3, doors=$4, updated_at=$5 WHERE id=$6;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Manufacturer, user.Design, user.Style, user.Doors, user.UpdatedAt, user.Id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUserById(id int) error {
	stmt, err := config.GetDb().Prepare("DELETE FROM users WHERE id=$1;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(id)
	return err
}
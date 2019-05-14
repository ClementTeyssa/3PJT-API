package models

import (
	"log"
	"time"

	"github.com/ClementTeyssa/3PJT-API/config"
)

//TODO: do comments for warning
type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	PublicKeyN string    `json:"publickeyn"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Users []User

func NewUser(user *User) {
	if user == nil {
		log.Fatal(user)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO users (email, password, publickeyn, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id;", user.Email, user.Password, user.PublicKeyN, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)

	if err != nil {
		log.Fatal(err)
	}
}

func FindUserById(id int) *User {
	var user User
	row := config.GetDb().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.PublicKeyN, &user.CreatedAt, &user.UpdatedAt)

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
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.PublicKeyN, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return &users
}

func UpdateUser(user *User) {
	user.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE users SET email=$1, password=$2, publickeyn=$3, updated_at=$4 WHERE id=$5;")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(user.Email, user.Password, user.PublicKeyN, user.UpdatedAt, user.ID)
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

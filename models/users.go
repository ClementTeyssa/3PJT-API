package models

import (
	"log"
	"time"

	"github.com/ClementTeyssa/3PJT-API/config"
)

//TODO: do comments for warning
type User struct {
	ID         int       `json:"id" validate:"omitempty,uuid"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Adress     string    `json:"adress" validate:"required"`
	PrivateKey []byte    `json:"privatekey" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Users []User

func UsersSize() int {
	rows, err := config.GetDb().Query("SELECT COUNT(*) as count FROM users")

	if err != nil {
		log.Panic(err)
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Panic(err)
		}
	}

	return count
}

func NewUser(user *User) {
	if user == nil {
		log.Panic(user)
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err := config.GetDb().QueryRow("INSERT INTO users (email, password, adress, privatekey, created_at, updated_at) VALUES ($1,$2,$3,$4,$5, $6) RETURNING id;", user.Email, user.Password, user.Adress, user.PrivateKey, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)

	if err != nil {
		log.Panic(err)
	}
}

func FindUserById(id int) *User {
	var user User
	row := config.GetDb().QueryRow("SELECT * FROM users WHERE id = $1;", id)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Adress, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Panic(err)
	}
	return &user
}

func FindUserByEmail(email string) *User {
	if UserWithEmailSize(email) > 0 {
		var user User
		row := config.GetDb().QueryRow("SELECT * FROM users WHERE email = $1;", email)
		err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Adress, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Panic(err)
		}
		return &user
	} else {
		log.Panic("Email doesn't exist")
		return nil
	}
}

func UserWithEmailSize(email string) int {
	rows, err := config.GetDb().Query("SELECT COUNT(*) as count FROM users WHERE email = $1;", email)

	if err != nil {
		log.Panic(err)
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Panic(err)
		}
	}

	return count
}

func AllUsers() *Users {
	var users Users
	rows, err := config.GetDb().Query("SELECT * FROM users")
	if err != nil {
		log.Panic(err)
	}
	// Close rows after all readed
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Adress, &user.PrivateKey, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		users = append(users, user)
	}
	return &users
}

func UpdateUser(user *User) {
	user.UpdatedAt = time.Now()
	stmt, err := config.GetDb().Prepare("UPDATE users SET email=$1, password=$2, adress=$3, privatekey=$4, updated_at=$5 WHERE id=$6;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(user.Email, user.Password, user.Adress, user.PrivateKey, user.UpdatedAt, user.ID)
	if err != nil {
		log.Panic(err)
	}
}

func DeleteUserById(id int) error {
	stmt, err := config.GetDb().Prepare("DELETE FROM users WHERE id=$1;")
	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(id)
	return err
}
